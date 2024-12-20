"use client";
import React, { ReactNode, useEffect } from 'react';
import { getUser } from '@/context/Auth';
import { useShowError } from '@/context/Error';

interface MasterLayoutProps {
    children: ReactNode | ReactNode[];
}

const MasterLayout: React.FC<MasterLayoutProps> = ({ children }) => {
    const User = getUser();
    const showError = useShowError();
        
    useEffect(() => {
        if (!User || User.Role !== 'efa') {
            window.location.href = '/';
            showError('You are not authorized to view this page');
        }
    }, [User]);

    return (
        <div>
            {User && User.Role === 'efa' ? children : null}
        </div>
    );
};

export default MasterLayout;