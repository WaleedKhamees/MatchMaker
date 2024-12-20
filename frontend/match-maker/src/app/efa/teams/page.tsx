"use client";
import { useShowError } from "@/context/Error";
import { Team, Stadium } from "@/types";
import { fetchCreateTeam, fetchStadiums, fetchTeams, fetchUpdateTeam } from "@/utils/api";
import Link from "next/link";
import React, { useState, useEffect } from "react";

const TeamsPage: React.FC = () => {
  const [teams, setTeams] = useState<Team[]>([]);
  const [stadiums, setStadiums] = useState<Stadium[]>([]);
  const [showForm, setShowForm] = useState(false);
  const [currentTeam, setCurrentTeam] = useState<Team | null>(null);
  const showError = useShowError();

  useEffect(() => {
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
    getTeams();
    getStadiums();
  }, []);

  const handleAddEditClick = (team?: Team) => {
    setCurrentTeam(team || null);
    setShowForm(true);
  };

  const handleFormClose = () => {
    setShowForm(false);
    setCurrentTeam(null);
  };

  const handleDelete = async (id: number) => {
    try {
      // await fetchDeleteTeam(id);
      // setTeams(teams.filter((team) => team.Id !== id));
    } catch (error: any) {
      // showError(error.message);
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
        <h1 className="text-3xl font-bold text-gray-800">Teams</h1>
        <button
          className="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700"
          onClick={() => handleAddEditClick()}
        >
          Add Team
        </button>
      </div>
      <ul className="space-y-6">
        {teams.map((team) => (
          <li
            key={team.Id}
            className="flex flex-col p-6 border rounded-lg shadow-lg bg-white space-y-4"
          >
            <div className="flex justify-between items-center">
              <div className="flex items-center space-x-6">
                {team.LogoUrl !== "" ? (
                  <img
                    src={team.LogoUrl}
                    alt={`${team.Name} logo`}
                    className="w-20 h-20 object-cover rounded-full"
                  />
                ) : null}
                <span className="text-xl font-semibold text-gray-800">
                  {team.Name}
                </span>
              </div>
              <div className="flex space-x-4">
                <button
                  className="bg-yellow-500 text-white px-4 py-2 rounded-lg hover:bg-yellow-600"
                  onClick={() => handleAddEditClick(team)}
                >
                  Edit
                </button>
                <button
                  className="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600"
                  onClick={() => handleDelete(team.Id!)}
                >
                  Delete
                </button>
              </div>
            </div>
            <div className="text-sm space-y-2 text-gray-700">
              <p>
                <strong>City:</strong> {team.City}
              </p>
              <p>
                <strong>Coach:</strong> {team.Coach}
              </p>
              <p>
                <strong>Founded:</strong> {team.Founded}
              </p>
            </div>
          </li>
        ))}
      </ul>
      {showForm && (
        <TeamForm
          team={currentTeam}
          onClose={handleFormClose}
          onSave={() => {}}
          addTeam={(newTeam) => setTeams([...teams, newTeam])}
          editTeam={(updatedTeam) =>
            setTeams(
              teams.map((t) => (t.Id === updatedTeam.Id ? updatedTeam : t))
            )
          }
          stadiums={stadiums}
        />
      )}
    </div>
  );
};

const TeamForm: React.FC<{
  team: Team | null;
  onClose: () => void;
  onSave: () => void;
  addTeam: (team: Team) => void;
  editTeam: (team: Team) => void;
  stadiums: Stadium[];
}> = ({ team, onClose, onSave, addTeam, editTeam, stadiums }) => {
  const [formState, setFormState] = useState<Team>(
    team || {
      Name: "",
      City: "",
      StadiumId: stadiums[0]?.Id || 0,
      Coach: "",
      Description: "",
      Founded: 0,
      LogoUrl: "",
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
    const ParsedFormState = {
      ...formState,
      StadiumId: Number(formState.StadiumId),
      Founded: Number(formState.Founded),
    };
    try {
      if (formState.Id) {
        const data = await fetchUpdateTeam(formState);
        editTeam(data);
      } else {
        const data = await fetchCreateTeam(ParsedFormState);
        addTeam(data);
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
        {team ? "Edit Team" : "Add Team"}
      </h2>
      <label className="block mb-4">
        Name:
        <input
          type="text"
          name="Name"
          value={formState.Name}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
        />
      </label>
      <label className="block mb-4">
        City:
        <input
          type="text"
          name="City"
          value={formState.City}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
        />
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
        Coach:
        <input
          type="text"
          name="Coach"
          value={formState.Coach}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
        />
      </label>
      <label className="block mb-4">
        Description:
        <input
          type="text"
          name="Description"
          value={formState.Description}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
        />
      </label>
      <label className="block mb-4">
        Founded:
        <input
          type="text"
          name="Founded"
          value={formState.Founded}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
        />
      </label>
      <label className="block mb-4">
        Logo URL:
        <input
          type="text"
          name="LogoUrl"
          value={formState.LogoUrl}
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

export default TeamsPage;
