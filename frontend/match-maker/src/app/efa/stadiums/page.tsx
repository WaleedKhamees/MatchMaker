"use client";
import { useShowError } from "@/context/Error";
import { Stadium } from "@/types";
import {
  fetchCreateStadium,
  fetchDeleteStadium,
  fetchStadiums,
  fetchUpdateStadium,
} from "@/utils/api";
import Link from "next/link";
import React, { useState, useEffect } from "react";

const StadiumsPage: React.FC = () => {
  const [stadiums, setStadiums] = useState<Stadium[]>([]);
  const [showForm, setShowForm] = useState(false);
  const [currentStadium, setCurrentStadium] = useState<Stadium | null>(null);
  const showError = useShowError();

  useEffect(() => {
    const getStadiums = async () => {
      const data = await fetchStadiums();
      setStadiums(data);
    };
    getStadiums();
  }, []);

  const handleAddEditClick = (stadium?: Stadium) => {
    setCurrentStadium(stadium || null);
    setShowForm(true);
  };

  const handleFormClose = () => {
    setShowForm(false);
    setCurrentStadium(null);
  };

  const handleDelete = async (id: number) => {
    try {
      await fetchDeleteStadium(id);
      setStadiums(stadiums.filter((stadium) => stadium.Id !== id));
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
        <h1 className="text-3xl font-bold text-gray-800">Stadiums</h1>
        <button
          className="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700"
          onClick={() => handleAddEditClick()}
        >
          Add Stadium
        </button>
      </div>
      <ul className="space-y-6">
        {stadiums.map((stadium) => (
          <li
            key={stadium.Id}
            className="flex flex-col p-6 border rounded-lg shadow-lg bg-white space-y-4"
          >
            <div className="flex justify-between items-center">
              <span className="text-xl font-semibold text-gray-800">
                {stadium.Name}
              </span>
              <div className="flex space-x-4">
                <button
                  className="bg-yellow-500 text-white px-4 py-2 rounded-lg hover:bg-yellow-600"
                  onClick={() => handleAddEditClick(stadium)}
                >
                  Edit
                </button>
                <button
                  className="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600"
                  onClick={() => handleDelete(stadium.Id!)}
                >
                  Delete
                </button>
              </div>
            </div>
            <div className="text-sm space-y-2 text-gray-700">
              <p>
                <strong>Capacity:</strong> {stadium.Capacity} seats
              </p>
              <p>
                <strong>VIP Rows:</strong> {stadium.VipRows}
              </p>
              <p>
                <strong>Seats Per Row:</strong> {stadium.SeatsPerRow}
              </p>
            </div>
          </li>
        ))}
      </ul>
      {showForm && (
        <StadiumForm
          stadium={currentStadium}
          onClose={handleFormClose}
          onSave={fetchStadiums}
          addStadium={(newStadium) => setStadiums([...stadiums, newStadium])}
          editStadium={(updatedStadium) =>
            setStadiums(
              stadiums.map((s) =>
                s.Id === updatedStadium.Id ? updatedStadium : s
              )
            )
          }
        />
      )}
    </div>
  );
};

const StadiumForm: React.FC<{
  stadium: Stadium | null;
  onClose: () => void;
  onSave: () => void;
  addStadium: (stadium: Stadium) => void;
  editStadium: (stadium: Stadium) => void;
}> = ({ stadium, onClose, onSave, addStadium, editStadium }) => {
  const [formState, setFormState] = useState<Stadium>(
    stadium || { Name: "", Capacity: 0, VipRows: 0, SeatsPerRow: 0 }
  );
  const showError = useShowError();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
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
      Capacity: Number(formState.Capacity),
      VipRows: Number(formState.VipRows),
      SeatsPerRow: Number(formState.SeatsPerRow),
    };
    try {
      if (formState.Id) {
        const data = await fetchUpdateStadium(parsedFormState);
        editStadium(data);
      } else {
        const data = await fetchCreateStadium(parsedFormState);
        addStadium(data);
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
        {stadium ? "Edit Stadium" : "Add Stadium"}
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
        Capacity:
        <input
          type="number"
          name="Capacity"
          value={formState.Capacity}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
        />
      </label>
      <label className="block mb-4">
        Vip Rows:
        <input
          type="number"
          name="VipRows"
          value={formState.VipRows}
          onChange={handleChange}
          className="border p-2 w-full rounded-lg"
        />
      </label>
      <label className="block mb-4">
        Seats Per Row:
        <input
          type="number"
          name="SeatsPerRow"
          value={formState.SeatsPerRow}
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

export default StadiumsPage;
