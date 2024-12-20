"use client";
import React, { useState, useEffect } from "react";
import { User } from "@/types";
import { useShowError } from "@/context/Error";
import Link from "next/link";
import { fetchApproveUser, fetchUnapprovedUsers } from "@/utils/api";

const App: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const showError = useShowError();

  // Fetch unapproved users from the backend using fetch API
  useEffect(() => {
    const getUnapprovedUsers = async () => {
      setLoading(true);
      try {
        const data = await fetchUnapprovedUsers();
        setUsers(data);
      } catch (error: any) {
        showError(error.message);
      } finally {
        setLoading(false);
      }
    };
    getUnapprovedUsers();
  }, []);

  // Approve or disapprove the user using fetch API
  const handleApproval = async (username: string, approved: boolean) => {
    try {
      const newUser = await fetchApproveUser(username, approved);

      setUsers((prevUsers) => {
        return prevUsers.filter((user) => user.Username !== username);
      });
    } catch (error: any) {
      showError(error.message);
    }
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="p-8">
      <Link
        href="/master"
        className="text-blue-500 hover:underline mb-4 inline-block"
      >
        Back to Master Page
      </Link>
      <h1 className="text-3xl font-bold mb-6 text-center">User Approval</h1>
      <table className="min-w-full table-auto border-collapse border border-gray-300">
        <thead>
          <tr>
            <th className="border p-2">Username</th>
            <th className="border p-2">Email</th>
            <th className="border p-2">Gender</th>
            <th className="border p-2">First Name</th>
            <th className="border p-2">Last Name</th>
            <th className="border p-2">Birthdate</th>
            <th className="border p-2">Role</th>
            <th className="border p-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          {users?.map((user) => (
            <tr key={user.Username}>
              <td className="border p-2">{user.Username}</td>
              <td className="border p-2">{user.Email}</td>
              <td className="border p-2">{user.Gender}</td>
              <td className="border p-2">{user.Firstname}</td>
              <td className="border p-2">{user.Lastname}</td>
              <td className="border p-2">
                {new Date(user.Birthdate).toLocaleDateString()}
              </td>
              <td className="border p-2">{user.Role}</td>
              <td className="border p-2">
                <button
                  className="bg-green-500 text-white p-2 rounded mr-2 hover:bg-green-600"
                  onClick={() => handleApproval(user.Username, true)}
                >
                  Approve
                </button>
                <button
                  className="bg-red-500 text-white p-2 rounded mr-2 hover:bg-red-600"
                  onClick={() => handleApproval(user.Username, false)}
                >
                  Disapprove
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default App;
