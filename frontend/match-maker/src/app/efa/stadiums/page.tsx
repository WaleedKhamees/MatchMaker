"use client";
import { useShowError } from "@/context/Error";
import { Stadium } from "@/types";
import {
  fetchCreateStadium,
  fetchDeleteStadium,
  fetchStadiums,
  fetchUpdateStadium,
} from "@/utils/api";
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
  }

  return (
    <div className="container mx-auto p-4">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-2xl font-bold">Stadiums</h1>
        <button
          className="bg-blue-500 text-white px-4 py-2 rounded"
          onClick={() => handleAddEditClick()}
        >
          Add Stadium
        </button>
      </div>
      <ul className="space-y-2">
        {stadiums.map((stadium) => (
          <li
            key={stadium.Id}
            className="flex flex-col p-4 border rounded space-y-2"
          >
            <div className="flex justify-between items-center">
              <span className="text-lg font-semibold">{stadium.Name}</span>
              <div className="flex space-x-2">
                <button
                  className="bg-yellow-500 text-white px-2 py-1 rounded"
                  onClick={() => handleAddEditClick(stadium)}
                >
                  Edit
                </button>
                <button
                  className="bg-red-500 text-white px-2 py-1 rounded"
                  onClick={() => handleDelete(stadium.Id!)}
                >
                  Delete
                </button>
              </div>
            </div>
            <div className="text-sm">
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
    <form onSubmit={handleSubmit} className="bg-white p-4 rounded shadow-md">
      <h2 className="text-xl font-bold mb-4">
        {stadium ? "Edit Stadium" : "Add Stadium"}
      </h2>
      <label className="block mb-2">
        Name:
        <input
          type="text"
          name="Name"
          value={formState.Name}
          onChange={handleChange}
          className="border p-2 w-full"
        />
      </label>
      <label className="block mb-2">
        Capacity:
        <input
          type="number"
          name="Capacity"
          value={formState.Capacity}
          onChange={handleChange}
          className="border p-2 w-full"
        />
      </label>
      <label className="block mb-2">
        Vip Rows:
        <input
          type="number"
          name="VipRows"
          value={formState.VipRows}
          onChange={handleChange}
          className="border p-2 w-full"
        />
      </label>
      <label className="block mb-2">
        Seats Per Row:
        <input
          type="number"
          name="SeatsPerRow"
          value={formState.SeatsPerRow}
          onChange={handleChange}
          className="border p-2 w-full"
        />
      </label>
      <div className="flex justify-end space-x-2">
        <button
          type="submit"
          className="bg-green-500 text-white px-4 py-2 rounded"
        >
          Save
        </button>
        <button
          type="button"
          onClick={onClose}
          className="bg-red-500 text-white px-4 py-2 rounded"
        >
          Cancel
        </button>
      </div>
    </form>
  );
};

export default StadiumsPage;
