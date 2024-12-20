"use client";
import { useShowError } from "@/context/Error";
import { Match, Team, Stadium, MatchRequest } from "@/types";
import {
  //   fetchCreateMatch,
  fetchMatches,
  fetchTeams,
  fetchStadiums,
  fetchCreateMatch,
  fetchDeleteMatch,
  //   fetchUpdateMatch,
} from "@/utils/api";
import Link from "next/link";
import React, { useState, useEffect } from "react";

const MatchesPage: React.FC = () => {
  const [matches, setMatches] = useState<Match[]>([]);
  const [teams, setTeams] = useState<Team[]>([]);
  const [stadiums, setStadiums] = useState<Stadium[]>([]);
  const [showForm, setShowForm] = useState(false);
  const [currentMatch, setCurrentMatch] = useState<Match | null>(null);
  const showError = useShowError();

  useEffect(() => {
    const getMatches = async () => {
      const data = await fetchMatches();
      setMatches(data?? []);
    };
    const getTeams = async () => {
      const data = await fetchTeams();
      setTeams(data);
    };
    const getStadiums = async () => {
      try {
        const data = await fetchStadiums();
        setStadiums(data);
      } catch (error: any) {
        showError(error.message);
      }
    };
    getMatches();
    getTeams();
    getStadiums();
  }, []);

  const handleAddEditClick = (match?: Match) => {
    setCurrentMatch(match || null);
    setShowForm(true);
  };

  const handleFormClose = () => {
    setShowForm(false);
    setCurrentMatch(null);
  };

  const handleDelete = async (id: number) => {
    try {
      await fetchDeleteMatch(id);
      setMatches(matches.filter((match) => match.Id !== id));
    } catch (error: any) {
      showError(error.message);
    }
  };

  return (
    <div className="container mx-auto p-6 bg-gray-100 min-h-screen">
      <Link
        href="/efa"
        className="text-blue-600 hover:underline mb-6 inline-block"
      >
        Back to EFA Page
      </Link>
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold text-gray-800">Matches</h1>
        <button
          className="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700"
          onClick={() => handleAddEditClick()}
        >
          Add Match
        </button>
      </div>
      <ul className="space-y-6">
        {matches?.map((match) => (
          <li
            key={match.Id}
            className="flex flex-col p-6 border rounded-lg shadow-lg bg-white space-y-4"
          >
            <div className="flex justify-between items-center">
              <div className="flex items-center space-x-6">
                <Link href={`/match/${match.Id}`}>
                  <span className="text-xl font-semibold text-gray-800">
                    {match.HomeTeam.Name} vs {match.AwayTeam.Name}
                  </span>
                </Link>
              </div>
              <div className="flex space-x-4">
                <button
                  className="bg-yellow-500 text-white px-4 py-2 rounded-lg hover:bg-yellow-600"
                  onClick={() => handleAddEditClick(match)}
                >
                  Edit
                </button>
                <button
                  className="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600"
                  onClick={() => handleDelete(match.Id!)}
                >
                  Delete
                </button>
              </div>
            </div>
            <div className="text-sm space-y-2 text-gray-700">
              <p>
                <strong>Date:</strong> {match.Date}
              </p>
              <p>
                <strong>Stadium: </strong>
                {match.Stadium.Name}
              </p>
              <p>
                <strong>Main Referee:</strong> {match.MainReferee}
              </p>
              <p>
                <strong>Line Man 1:</strong> {match.Lineman1}
              </p>
              <p>
                <strong>Line Man 2:</strong> {match.Lineman2}
              </p>
            </div>
          </li>
        ))}
      </ul>
      {showForm && (
        <MatchForm
          match={currentMatch}
          onClose={handleFormClose}
          onSave={() => {}}
          addMatch={(newMatch) => matches? setMatches([...matches, newMatch]) : setMatches([newMatch])}
          editMatch={(updatedMatch) =>
            setMatches(
              matches.map((m) => (m.Id === updatedMatch.Id ? updatedMatch : m))
            )
          }
          teams={teams}
          stadiums={stadiums}
        />
      )}
    </div>
  );
};

const MatchForm: React.FC<{
  match: Match | null;
  onClose: () => void;
  onSave: () => void;
  addMatch: (match: Match) => void;
  editMatch: (match: Match) => void;
  teams: Team[];
  stadiums: Stadium[];
}> = ({ match, onClose, onSave, teams, stadiums, addMatch }) => {
  const [formState, setFormState] = useState<MatchRequest>(
    {
      HomeTeamId: teams[0].Id || 0,
      AwayTeamId: teams[0].Id || 0,
      StadiumId: stadiums[0]?.Id || 0,
      Date: "",
      MainReferee: "",
      Lineman1: "",
      Lineman2: "",
    }
  );
  const showError = useShowError();

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    const { name, value } = e.target;
    setFormState((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const parsedFormState = {
        ...formState,
        HomeTeamId: Number(formState.HomeTeamId),
        AwayTeamId: Number(formState.AwayTeamId),
        StadiumId: Number(formState.StadiumId),
        Date: new Date(formState.Date).toISOString(),
    };
    console.log(parsedFormState);
    try {
      if (formState.Id) {
        // const data = await fetchUpdateMatch(formState);
        // editMatch(data);
      } else {
        const data = await fetchCreateMatch(parsedFormState);
        addMatch(data);
      }
      onSave();
      onClose();
    } catch (error: any) {
      showError(error.message);
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="bg-white p-8 rounded-lg shadow-lg mx-auto mt-4"
    >
      <h2 className="text-2xl font-bold mb-6 text-gray-800">
        {match ? "Edit Match" : "Add Match"}
      </h2>
      <label className="block mb-4">
        Home Team:
        <select
          name="HomeTeamId"
          value={formState.HomeTeamId}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
          required
        >
          <option value="" disabled>
            Select a home team
          </option>
          {teams.map((team) => (
            <option key={team.Id} value={team.Id}>
              {team.Name}
            </option>
          ))}
        </select>
      </label>
      <label className="block mb-4">
        Away Team:
        <select
          name="AwayTeamId"
          value={formState.AwayTeamId}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
          required
        >
          <option value="" disabled>
            Select an away team
          </option>
          {teams.map((team) => (
            <option key={team.Id} value={team.Id}>
              {team.Name}
            </option>
          ))}
        </select>
      </label>
      <label className="block mb-4">
        Stadium:
        <select
          name="StadiumId"
          value={formState.StadiumId}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
          required
        >
          <option value="" disabled>
            Select a stadium
          </option>
          {stadiums.map((stadium) => (
            <option key={stadium.Id} value={stadium.Id}>
              {stadium.Name}
            </option>
          ))}
        </select>
      </label>
      <label className="block mb-4">
        Date:
        <input
          type="datetime-local"
          name="Date"
          value={formState.Date}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
        />
      </label>
      <label className="block mb-4">
        Main Referee:
        <input
          type="text"
          name="MainReferee"
          value={formState.MainReferee}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
        />
      </label>
      <label className="block mb-4">
        Line Man 1:
        <input
          type="text"
          name="Lineman1"
          value={formState.Lineman1}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
        />
      </label>
      <label className="block mb-4">
        Line Man 2:
        <input
          type="text"
          name="Lineman2"
          value={formState.Lineman2}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
        />
      </label>
      <div className="flex justify-end space-x-4">
        <button
          type="submit"
          className="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700"
        >
          Save
        </button>
        <button
          type="button"
          onClick={onClose}
          className="bg-red-600 text-white px-6 py-2 rounded-lg hover:bg-red-700"
        >
          Cancel
        </button>
      </div>
    </form>
  );
};

export default MatchesPage;
