"use client";

import { useEffect, useState } from "react";
import { Match, Stadium } from "@/types";
import { fetchMatchById, fetchMatchSeats } from "@/utils/api";
import io from "socket.io-client";
import { useParams } from "next/navigation";
import { GetAuthToken, getUser } from "@/context/Auth";
import { Socket } from "socket.io-client";
import { useShowError } from "@/context/Error";

type reservedSeat = {
  Username: string;
  SeatRow: number;
  SeatColumn: number;
};

const MatchPage: React.FC = () => {
  const { id } = useParams();

  const [match, setMatch] = useState<Match | null>(null);
  const [reservedSeats, setReservedSeats] = useState<reservedSeat[]>([]);
  const [socket, setSocket] = useState<Socket | null>(null);
  const showError = useShowError();

  useEffect(() => {
    setSocket(
      io("http://localhost:8000/socket/match", {
        transports: ["websocket"],
      })
    );

    return () => {
      if (socket) {
        socket.disconnect();
      }
    };
  }, []);

  useEffect(() => {
    if (!socket) {
      return;
    }

    socket.on("onreserve", (msg: string) => {
      const data = JSON.parse(msg);

      if (data.typeofreq === "reserve") {
        setReservedSeats((seats) => [
          ...seats,
          {
            Username: data.username,
            SeatRow: data.seatrow,
            SeatColumn: data.seatcol,
          },
        ]);
      }
      else if (data.typeofreq === "cancel") {
        setReservedSeats((seats) =>
          seats.filter(
            (seat) => !(seat.SeatRow === data.seatrow && seat.SeatColumn === data.seatcol)
          )
        );
      }
    });

    if (id) {
      const getMatch = async () => {
        const data = await fetchMatchById(Number(id));
        setMatch(data);
        const payload = JSON.stringify({ matchid: Number(id) });
        socket.emit("subscribe", payload);
        try {
          const seats = await fetchMatchSeats(Number(id));
          console.log("Seats:", seats);
          setReservedSeats(seats);
        } catch (error: any) {
          showError(error.message);
        }
      };
      getMatch();
    }

    return () => {
      if (socket) {
        socket.off("onreserve");
      }
    };
  }, [socket, id]);

  if (!match) {
    return <div>Loading...</div>;
  }

  const { Stadium } = match;

  const handleSeatClick = (row: number, col: number) => {
    socket!.emit(
      "reserve",
      JSON.stringify({
        token: GetAuthToken(),
        matchid: Number(id),
        seatrow: row,
        seatcol: col,
      })
    );

    setReservedSeats((seats) => [
      ...seats,
      {
        Username: getUser()?.Username!,
        SeatRow: row,
        SeatColumn: col,
      },
    ]);

  };

  const handleCancelSeat = (row: number, col: number) => {
    socket!.emit(
      "cancel",
      JSON.stringify({
        token: GetAuthToken(),
        matchid: Number(id),
        seatrow: row,
        seatcol: col,
      })
    );

    setReservedSeats((seats) =>
      seats.filter(
        (seat) => !(seat.SeatRow === row && seat.SeatColumn === col)
      )
    );
  }

  return (
    <div className="container mx-auto p-6 bg-gray-100 min-h-screen">
      <h1 className="text-3xl font-bold text-gray-800 mb-6">
        {match.HomeTeam.Name} vs {match.AwayTeam.Name}
      </h1>
      <div className="flex justify-center items-center w-full">
        <SeatGrid
          Stadium={Stadium}
          reservedSeats={reservedSeats}
          handleSeatClick={handleSeatClick}
          handleCancelSeat={handleCancelSeat}
        />
      </div>
    </div>
  );
};

type SeatGridProps = {
  Stadium: Stadium;
  reservedSeats: reservedSeat[];
  handleSeatClick: (row: number, col: number) => void;
  handleCancelSeat: (row: number, col: number) => void;
};

function SeatGrid({ Stadium, reservedSeats, handleSeatClick, handleCancelSeat }: SeatGridProps) {
  return (
    <div className="flex flex-col">
      {Array.from({ length: Stadium.VipRows }).map((_, row: number) => (
        <div key={row} className="flex">
          {Array.from({ length: Stadium.SeatsPerRow }).map((_, col: number) => {
            const reserved = Array.isArray(reservedSeats) && reservedSeats.find(
              (seat) => seat.SeatRow === row && seat.SeatColumn === col
            );
            const seatNumber = row * Stadium.SeatsPerRow + col + 1;

            if (reserved && reserved.Username === getUser()?.Username) {
              return (
                <button
                  key={col}
                  onClick={() => handleCancelSeat(row, col)}
                  className="bg-white-500 border font-bold p-2 w-12 h-12"
                >
                  {seatNumber}
                </button>
              );
            }

            if (reserved) {
              return (
                <button
                  key={col}
                  className="bg-red-500 border font-bold p-2 w-12 h-12"
                >
                  {seatNumber}
                </button>
              );
            }

            return (
              <button
                key={col}
                onClick={() => handleSeatClick(row, col)}
                className="bg-green-500 border font-bold p-2 w-12 h-12"
              >
                {seatNumber}
              </button>
            );
          })}
        </div>
      ))}
    </div>
  );
}

export default MatchPage;
