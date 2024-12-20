"use client";
import React from "react";
import Link from "next/link";

const AdminPage = () => {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <h1 className="text-4xl font-bold mb-4">Welcome to the Admin Page</h1>
      <p className="text-lg mb-6">Please go to the manage events page.</p>
      <div className="flex items-center gap-4">
        <Link
          href="/efa/matches"
          className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-700 transition duration-300"
        >
          Manage Events
        </Link>
        <Link
          href="/efa/stadiums"
          className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-700 transition duration-300"
        >
          Manage Stadiums
        </Link>
        <Link
          href="/efa/teams"
          className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-700 transition duration-300"
        >
          Manage Teams
        </Link>
      </div>
    </div>
  );
};

export default AdminPage;
