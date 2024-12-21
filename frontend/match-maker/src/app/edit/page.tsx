"use client";

import { useState } from "react";
import { fetchEditUser } from "@/utils/api";
import { useShowError } from "@/context/Error";
import { getUser } from "@/context/Auth";
import Link from "next/link";

const EditProfilePage = () => {
  const showError = useShowError();
  const user = getUser();

  if (!user) {
    return (
      <p>
        You must be logged in to access this page.{" "}
        <Link href="/" className="text-blue-500">
          Click here to go back to the home page.
        </Link>
      </p>
    );
  }

  const [formData, setFormData] = useState({
    username: user.Username,
    firstname: user.Firstname,
    lastname: user.Lastname,
    email: user.Email,
    password: "",
    confirmPassword: "",
    gender: user.Gender,
    birthdate: user.Birthdate,
    city: user.City,
    address: user.Address,
    role: user.Role,
  });

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (formData.password !== formData.confirmPassword) {
      showError("Passwords do not match");
      return;
    }

    if (formData.role === "") {
      showError("Please select a role");
      return;
    }

    if (formData.gender === "") {
      showError("Please select a Gender");
      return;
    }

    if (formData.birthdate === "") {
      showError("Please select a Birthdate");
      return;
    }

    formData.birthdate = new Date(formData.birthdate).toISOString();

    try {
      const { message } = await fetchEditUser(formData);
      showError("Profile updated successfully", true);
    } catch (error: any) {
      showError(error.message);
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <div className="bg-white p-8 rounded-lg w-96 max-h-[90vh] overflow-y-auto">
        <h2 className="text-2xl font-bold text-gray-800 mb-6">Edit Profile</h2>
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
            disabled
            className="w-full p-2 border rounded bg-gray-300"
          />
          <input
            type="email"
            name="email"
            placeholder="Email"
            value={formData.email}
            disabled
            className="w-full p-2 border rounded bg-gray-300"
          />
          <input
            type="password"
            name="password"
            placeholder="Password"
            value={formData.password}
            onChange={handleChange}
            className="w-full p-2 border rounded"
          />
          <input
            type="password"
            name="confirmPassword"
            placeholder="Confirm Password"
            value={formData.confirmPassword}
            onChange={handleChange}
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
          <select
            name="role"
            value={formData.role}
            onChange={handleChange}
            required
            className="w-full p-2 border rounded "
          >
            <option value="">Select Role</option>
            <option value="customer">Customer</option>
            <option value="efa">EFA Admin</option>
            <option value="master">Master</option>
          </select>
          <input
            type="date"
            name="birthdate"
            value={formData.birthdate ? formData.birthdate.split("T")[0] : ""} // Ensure it's in YYYY-MM-DD format
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
            className="w-full p-2 border rounded"
          />
          <button
            type="submit"
            className="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600"
          >
            Save Changes
          </button>
        </form>
      </div>
    </div>
  );
};

export default EditProfilePage;
