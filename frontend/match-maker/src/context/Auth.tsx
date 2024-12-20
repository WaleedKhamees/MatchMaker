"use client";
import React, { createContext, ReactNode, useEffect, useState } from "react";
import { User } from "@/types";
import { useShowError } from "./Error";
import { fetchLogin } from "@/utils/api";

interface AuthContextProps {
  user: User | null;
  authToken: string | null;
  isLoggedIn: boolean;
  Login: (identifier: string, authToken: string) => void;
  Logout: () => void;
}

export const AuthContext = createContext<AuthContextProps>({
  user: null,
  authToken: null,
  isLoggedIn: false,
  Login: () => {},
  Logout: () => {},
});

export const AuthProvider = ({
  children,
}: {
  children: ReactNode | ReactNode[];
}) => {
  const [user, _setUser] = useState<User | null>(null);
  const [authToken, setAuthToken] = useState<string | null>(null);
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const showError = useShowError();

  function setUser(user: User | null) {
    if (!user) {
      localStorage.removeItem("user");
      return;
    }

    _setUser(user);
    localStorage.setItem("user", JSON.stringify(user));
  }

  function setToken(authToken: string | null) {
    const localStorageAuth = localStorage.getItem("authToken");

    if (!localStorageAuth) {
      localStorage.removeItem("authToken");
    }

    setAuthToken(authToken);
    if (authToken) {
      setAuthToken(authToken);
      localStorage.setItem("authToken", authToken);
    }
  }

  const Login = async (identifier: string, password: string) => {
    try {
      const { token, user } = await fetchLogin({ identifier, password });
      setUser(user);
      setIsLoggedIn(true);
      setToken(token);
      window.location.href = "/master"

    } catch (error: any) {
      showError(error.message);
    }
  };

  const Logout = () => {
    setUser(null);
    setToken(null);
    setIsLoggedIn(false);
    localStorage.removeItem("authToken");
    localStorage.removeItem("email");
    localStorage.removeItem("username");
    
    window.location.href = "/";
  };

  useEffect(() => {
    if (localStorage.getItem("authToken")) {
      setAuthToken(localStorage.getItem("authToken"));
      setIsLoggedIn(true);
    }
    const userString = localStorage.getItem("user");
    if (userString) {
      const user = JSON.parse(userString);
      setUser(user);
      if (user.Role === "master" && !window.location.pathname.includes("/master")) { 
        window.location.href = "/master";
      }
      else if (user.Role === "efa" && !window.location.pathname.includes("/efa")) {
        window.location.href = "/efa";
      }
    }
  }, []);


  return (
    <AuthContext.Provider
      value={{
        user,
        authToken,
        isLoggedIn,
        Login,
        Logout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export const GetAuthToken = () => {
  const authToken = localStorage.getItem("authToken");
  return authToken;
};

export const getUser = () => {
  const userString = localStorage.getItem("user");
  if (userString) {
    return JSON.parse(userString);
  }
  return null;
}