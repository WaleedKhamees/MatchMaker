import React, { useContext } from "react";
import { clearError, ErrorModalContext } from "../context/Error";

const ErrorModal: React.FC = () => {
  const { error } = useContext(ErrorModalContext);

  if (!error) {
    return null;
  }

  return (
    <div className="fixed inset-0 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg p-8 shadow-lg">
        <h2 className="text-xl font-bold mb-4">Error</h2>
        <p className="text-red-500 mb-4">{error}</p>
        <button
          className="bg-red-500 text-white px-4 py-2 rounded-lg"
          onClick={clearError}
        >
          ✕
        </button>
      </div>
    </div>
  );
};

export default ErrorModal;
