import React, { createContext, ReactNode, useContext, useState } from 'react';

export const ErrorModalContext = createContext<{
    error: string | null;
    setError: React.Dispatch<React.SetStateAction<string | null>>;
}>({
    error: null,
    setError: () => {},
});

export const ErrorModalProvider = ({ children } : { children: ReactNode | ReactNode[] }) => {
    const [error, setError] = useState<string | null>(null);

    return (
        <ErrorModalContext.Provider value={{ error, setError }}>
            {children}
        </ErrorModalContext.Provider>
    );
};

export const showError = (error: string) => {
    const { setError } = useContext(ErrorModalContext);
    setError(error);
}

export const clearError = () => {
    const { setError } = useContext(ErrorModalContext);
    setError(null);
}
