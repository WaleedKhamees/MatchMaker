"use client";
import React, { useContext } from "react";
import { clearError, ErrorModalContext } from "../../context/Error";

const ErrorModal: React.FC = () => {
  const { error, setError } = useContext(ErrorModalContext);

  if (!error) {
    return null;
  }

  return (
    <div className="fixed inset-0 flex items-center justify-center z-50 bg-black bg-opacity-80">
      <div className="bg-white rounded-lg p-8 shadow-2xl w-1/3 max-w-xl">
        <div className="flex justify-between items-center pb-6">
          <h2 className="text-2xl font-bold">Error</h2>
          <button
            className="bg-red-500 text-white px-4 py-2 rounded-lg"
            onClick={() => setError("")}
          >
            ✕
          </button>
        </div>
        <p className="text-red-500 mb-4 text-lg">{error}</p>
      </div>
    </div>
  );
};

export default ErrorModal;
