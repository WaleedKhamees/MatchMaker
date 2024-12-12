"use client";

import Link from "next/link";
import { useContext, useState } from "react";
import { fetchLogin } from "../utils/api";
import { AuthContext } from "../context/Auth";

function LoginModal({
  onClose,
  onLogin,
}: {
  onClose: () => void;
  onLogin: (loginData: { identifier: string; password: string }) => void;
}) {
  const [loginData, setLoginData] = useState({
    identifier: "",
    password: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setLoginData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    onLogin(loginData);
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-50">
      <div className="bg-white p-8 rounded-lg w-96">
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-bold text-gray-800">Login</h2>
          <button
            onClick={onClose}
            className="text-gray-600 hover:text-gray-900"
          >
            ✕
          </button>
        </div>

        <form onSubmit={handleSubmit} className="space-y-4">
          <input
            type="text"
            name="username"
            placeholder="Username or Email"
            value={loginData.identifier}
            onChange={handleChange}
            required
            className="w-full p-2 border rounded"
          />
          <input
            type="password"
            name="password"
            placeholder="Password"
            value={loginData.password}
            onChange={handleChange}
            required
            className="w-full p-2 border rounded"
          />

          <button
            type="submit"
            className="w-full bg-green-500 text-white p-2 rounded hover:bg-green-600"
          >
            Login
          </button>
        </form>
      </div>
    </div>
  );
}

const NavBar = () => {
  const [isLoginModalOpen, setIsLoginModalOpen] = useState(false);
  const { Login } = useContext(AuthContext);

  const handleLogin = async (loginData: {
    identifier: string;
    password: string;
  }) => {
    try {
      const response = await fetchLogin(loginData);

      if (response.ok) {
        const userData = (await response.json()) as { token: string };
        const { token } = userData;
        Login(loginData.identifier, token);
        setIsLoginModalOpen(false);
      } else {
        const errorText = await response.text();
        alert(`Login failed: ${errorText}`);
      }
    } catch (error) {
      console.error("Login error:", error);
      alert("An error occurred during login");
    }
  };

  return (
        <div className="flex justify-between items-center mb-10">
          <h1 className="text-4xl font-bold text-center text-gray-800">
            Match Maker: Upcoming Football Matches
          </h1>
          <div className="flex items-center space-x-4">
            {isLoggedIn ? (
              <>
                <span className="text-gray-700">Welcome, {currentUser}</span>
                <button 
                  onClick={handleLogout}
                  className="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded"
                >
                  Logout
                </button>
              </>
            ) : (
              <>
                <button 
                  onClick={() => setIsLoginModalOpen(true)}
                  className="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded"
                >
                  Login
                </button>
                <button 
                  onClick={() => setIsSignUpModalOpen(true)}
                  className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
                >
                  Sign Up
                </button>
              </>
            )}
          </div>
        </div>

  );
};

export default NavBar;
