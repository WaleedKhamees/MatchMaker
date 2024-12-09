"use client";

import React, { useState } from 'react';
import { Match } from '@/types/types';

export default function Home() {
  const [isSignUpModalOpen, setIsSignUpModalOpen] = useState(false);
  const [isLoginModalOpen, setIsLoginModalOpen] = useState(false);
  const [matches, setMatches] = useState<Match[]>([]);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [currentUser, setCurrentUser] = useState<string | null>(null);

  React.useEffect(() => {
    async function fetchMatches() {
      try {
        const response = await fetch('http://localhost:8080/matches', { 
          cache: 'no-store'
        });

        if (!response.ok) {
          throw new Error('Failed to fetch matches');
        }

        const data = await response.json();
        setMatches(data);
      } catch (error) {
        console.error('Error fetching matches:', error);
      }
    }

    fetchMatches();
  }, []);

  const handleLogin = async (loginData: { username: string, password: string }) => {
    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(loginData)
      });

      if (response.ok) {
        const userData = await response.json();
        setIsLoggedIn(true);
        setCurrentUser(userData.username);
        setIsLoginModalOpen(false);
        alert('Login successful!');
      } else {
        const errorText = await response.text();
        alert(`Login failed: ${errorText}`);
      }
    } catch (error) {
      console.error('Login error:', error);
      alert('An error occurred during login');
    }
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    setCurrentUser(null);
  };

  return (
    <div className="min-h-screen bg-gray-100 py-10">
      <div className="container mx-auto px-4">
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

        {matches.length === 0 ? (
          <div className="text-center text-gray-600">
            <p className="text-2xl">No matches scheduled at the moment</p>
          </div>
        ) : (
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
            {matches.map((match) => (
              <div 
                key={match.Id} 
                className="bg-white rounded-lg shadow-md overflow-hidden transform transition duration-300 hover:scale-105"
              >
                <div className="p-6">
                  <div className="flex justify-between items-center mb-4">
                    <h2 className="text-2xl font-semibold text-gray-800">Match #{match.Id}</h2>
                    <span className="bg-blue-100 text-blue-800 text-xs font-medium px-2.5 py-0.5 rounded">
                      Scheduled
                    </span>
                  </div>
                  
                  <div className="space-y-3 text-gray-600">
                    <div className="flex justify-between">
                      <span className="font-medium">Home Team:</span>
                      <span>Team ID {match.HomeTeamId}</span>
                    </div>
                    <div className="flex justify-between">
                      <span className="font-medium">Away Team:</span>
                      <span>Team ID {match.AwayTeamId}</span>
                    </div>
                    <div className="flex justify-between">
                      <span className="font-medium">Stadium:</span>
                      <span>Stadium ID {match.StadiumId}</span>
                    </div>
                    <div className="flex justify-between">
                      <span className="font-medium">Date:</span>
                      <span>{new Date(match.Date).toLocaleString()}</span>
                    </div>
                    <div className="border-t pt-3">
                      <h3 className="text-lg font-semibold mb-2 text-gray-700">Match Officials</h3>
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
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}

        {isSignUpModalOpen && (
          <SignUpModal onClose={() => setIsSignUpModalOpen(false)} />
        )}

        {isLoginModalOpen && (
          <LoginModal 
            onClose={() => setIsLoginModalOpen(false)}
            onLogin={handleLogin}
          />
        )}
      </div>
    </div>
  );
}

function SignUpModal({ onClose }: { onClose: () => void }) {
  const [formData, setFormData] = useState({
    username: '',
    firstname: '',
    lastname: '',
    email: '',
    password: '',
    confirmPassword: '',
    gender: '',
    birthdate: '',
    city: '',
    address: '',
    role: 'customer'
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    // Basic validation
    if (formData.password !== formData.confirmPassword) {
      alert("Passwords do not match");
      return;
    }

    try {
      const response = await fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ...formData,
          birthdate: new Date(formData.birthdate).toISOString()
        })
      });

      if (response.ok) {
        alert('Registration successful!');
        onClose();
      } else {
        const errorText = await response.text();
        alert(`Registration failed: ${errorText}`);
      }
    } catch (error) {
      console.error('Registration error:', error);
      alert('An error occurred during registration');
    }
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-50">
      <div className="bg-white p-8 rounded-lg w-96 max-h-[90vh] overflow-y-auto">
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-bold text-gray-800">Sign Up</h2>
          <button 
            onClick={onClose} 
            className="text-gray-600 hover:text-gray-900"
          >
            ✕
          </button>
        </div>
        
        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="grid grid-cols-2 gap-4">
            <input
              type="text"
              name="firstname"
              placeholder="First Name"
              value={formData.firstname}
              onChange={handleChange}
              required
              className="w-full p-2 border rounded"
            />
            <input
              type="text"
              name="lastname"
              placeholder="Last Name"
              value={formData.lastname}
              onChange={handleChange}
              required
              className="w-full p-2 border rounded"
            />
          </div>
          <input
            type="text"
            name="username"
            placeholder="Username"
            value={formData.username}
            onChange={handleChange}
            required
            className="w-full p-2 border rounded"
          />
          <input
            type="email"
            name="email"
            placeholder="Email"
            value={formData.email}
            onChange={handleChange}
            required
            className="w-full p-2 border rounded"
          />
          <input
            type="password"
            name="password"
            placeholder="Password"
            value={formData.password}
            onChange={handleChange}
            required
            className="w-full p-2 border rounded"
          />
          <input
            type="password"
            name="confirmPassword"
            placeholder="Confirm Password"
            value={formData.confirmPassword}
            onChange={handleChange}
            required
            className="w-full p-2 border rounded"
          />
          <select
            name="gender"
            value={formData.gender}
            onChange={handleChange}
            required
            className="w-full p-2 border rounded"
          >
            <option value="">Select Gender</option>
            <option value="male">Male</option>
            <option value="female">Female</option>
          </select>
          <input
            type="date"
            name="birthdate"
            value={formData.birthdate}
            onChange={handleChange}
            required
            className="w-full p-2 border rounded"
          />
          <input
            type="text"
            name="city"
            placeholder="City"
            value={formData.city}
            onChange={handleChange}
            required
            className="w-full p-2 border rounded"
          />
          <input
            type="text"
            name="address"
            placeholder="Address"
            value={formData.address}
            onChange={handleChange}
            required
            className="w-full p-2 border rounded"
          />
          
          <button 
            type="submit" 
            className="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600"
          >
            Sign Up
          </button>
        </form>
      </div>
    </div>
  );
}

function LoginModal({ 
  onClose, 
  onLogin 
}: { 
  onClose: () => void, 
  onLogin: (loginData: { username: string, password: string }) => void 
}) {
  const [loginData, setLoginData] = useState({
    username: '',
    password: ''
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setLoginData(prev => ({
      ...prev,
      [name]: value
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
            value={loginData.username}
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