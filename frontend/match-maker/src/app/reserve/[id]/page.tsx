"use client";
import { useRouter, useParams } from "next/navigation";
import { useEffect, useState } from "react";
import { Typography, Box, Button } from "@mui/material";

export default function ReservePage() {
  const router = useRouter();
  const { id } = useParams();
  const [stadiumId, setStadiumId] = useState<number | null>(null); // To store the stadium ID
  const [data, setData] = useState<any>(null);

  useEffect(() => {
    // Extract stadiumId from URL query parameters
    const urlParams = new URLSearchParams(window.location.search);
    const stadiumIdFromUrl = urlParams.get("stadiumId");

    if (stadiumIdFromUrl) {
      setStadiumId(parseInt(stadiumIdFromUrl));
    }

    const fetchData = async () => {
      try {
        const response = await fetch(`http://localhost:8000/match/${id}`);
        if (!response.ok) {
          throw new Error("Failed to fetch data");
        }
        const result = await response.json();
        setData(result);
      } catch (err) {
        console.error(err);
      }
    };

    fetchData();
  }, [id]);

  if (!data) return <div>Loading...</div>;

  return (
    <Box sx={{ padding: 3 }}>
      <Typography variant="h4" gutterBottom>
        Reserve Seat for Match: {data.HomeTeam.Name} vs {data.AwayTeam.Name}
      </Typography>
      <Typography variant="h6">
        Stadium ID: {stadiumId} {/* Display the stadium ID */}
      </Typography>

      {/* You can add your seat reservation form here */}
      <Button variant="contained" color="primary">
        Confirm Reservation
      </Button>
    </Box>
  );
}
