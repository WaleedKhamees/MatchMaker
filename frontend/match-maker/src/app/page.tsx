"use client";

import React, { useState } from "react";
import Link from "next/link";
import { fetchMatches } from "@/utils/api";
import { useShowError } from "@/context/Error";
import { Match } from "@/types";

export default function Home() {
  const [matches, setMatches] = useState<Match[]>([]);
  const showError = useShowError();

  React.useEffect(() => {
    async function getMatches() {
      try {
        const data = await fetchMatches();
        setMatches(data);
        console.log(data);
      } catch (error: any) {
        showError(error.message);
      }
    }
    getMatches();
  }, []);

  return (
    <div className="py-10">
      <div className="container mx-auto px-4">
        {!matches || matches.length === 0 ? (
          <div className="text-center">
            <p className="text-2xl">No matches scheduled at the moment</p>
          </div>
        ) : (
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
            {matches?.map((match) => (
              <div
                key={match.Id}
                className="bg-white rounded-lg shadow-md overflow-hidden transform transition duration-300 hover:scale-105 cursor-pointer"
              >
                <div className="p-6">
                  <div className="flex justify-between items-center mb-4">
                    <h2 className="text-2xl font-semibold text-gray-800">
                      Match #{match.Id}
                    </h2>
                    <span className="bg-blue-100 text-blue-800 text-xs font-medium px-2.5 py-0.5 rounded">
                      Scheduled
                    </span>
                  </div>

                  <div className="space-y-3 text-gray-600">
                    <div className="flex justify-between items-center">
                      <span className="font-medium">Home Team:</span>
                      <span>{match.HomeTeam.Name}</span>
                      <img src={match.HomeTeam.LogoUrl} alt={`${match.HomeTeam.Name} logo`} className="w-8 h-8"/>
                    </div>
                    <div className="flex justify-between items-center">
                      <span className="font-medium">Away Team:</span>
                      <span>{match.AwayTeam.Name}</span>
                      <img src={match.AwayTeam.LogoUrl} alt={`${match.AwayTeam.Name} logo`} className="w-8 h-8"/>
                    </div>
                    <div className="flex justify-between">
                      <span className="font-medium">Stadium:</span>
                      <span>{match.Stadium.Name}</span>
                    </div>
                    <div className="flex justify-between">
                      <span className="font-medium">Date:</span>
                      <span>{new Date(match.Date).toLocaleString()}</span>
                    </div>
                    <div className="border-t pt-3">
                      <h3 className="text-lg font-semibold mb-2 text-gray-700">
                        Match Officials
                      </h3>
                      <div className="flex justify-between">
                        <span className="font-medium">Main Referee:</span>
                        <span>{match.MainReferee}</span>
                      </div>
                      <div className="flex justify-between">
                        <span className="font-medium">Lineman 1:</span>
                        <span>{match.Lineman1}</span>
                      </div>
                      <div className="flex justify-between">
                        <span className="font-medium">Lineman 2:</span>
                        <span>{match.Lineman2}</span>
                      </div>
                    </div>
                    <div className="border-t pt-3">
                      <h3 className="text-lg font-semibold mb-2 text-gray-700">
                        Additional Details
                      </h3>
                      <div className="flex justify-between">
                        <span className="font-medium">Home Team Coach:</span>
                        <span>{match.HomeTeam.Coach}</span>
                      </div>
                      <div className="flex justify-between">
                        <span className="font-medium">Away Team Coach:</span>
                        <span>{match.AwayTeam.Coach}</span>
                      </div>
                      <div className="flex justify-between">
                        <span className="font-medium">Stadium Capacity:</span>
                        <span>{match.Stadium.Capacity}</span>
                      </div>
                    </div>
                  </div>
                  <div className="mt-4">
                    <Link
                      href={`/match/${match.Id}`}
                      passHref
                      className="inline-block bg-blue-500 text-white text-center py-2 px-4 rounded hover:bg-blue-600 transition duration-300"
                    >
                      View Match Details
                    </Link>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
