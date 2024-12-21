"use client";
import React, { createContext, ReactNode, useContext, useState } from 'react';

export const ErrorModalContext = createContext<{
    error: string | null;
    info: boolean;
    setError: React.Dispatch<React.SetStateAction<string | null>>;
    setInfo: React.Dispatch<React.SetStateAction<boolean>>;
}>({
    error: null,
    info: false,
    setError: () => {},
    setInfo: () => {},
});

export const ErrorModalProvider = ({ children } : { children: ReactNode | ReactNode[] }) => {
    const [error, setError] = useState<string | null>(null);
    const [info, setInfo] = useState<boolean>(false);

    return (
        <ErrorModalContext.Provider value={{ error, setError, setInfo, info }}>
            {children}
        </ErrorModalContext.Provider>
    );
};

export const useShowError = () => {
    const { setError, setInfo } = useContext(ErrorModalContext);
    return (error: string, info = false) => {
        setError(error);
        setInfo(info);
    };
};

export const clearError = () => {
    const { setError, setInfo } = useContext(ErrorModalContext);
    setError(null);
    setInfo(false);
}
