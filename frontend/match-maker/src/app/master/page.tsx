"use client";
import React from 'react';
import { useRouter } from 'next/router';
import Link from 'next/link';

const MasterPage = () => {
    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
            <h1 className="text-4xl font-bold mb-4">Welcome to the Master Page</h1>
            <p className="text-lg mb-6">Please go to the approve user page.</p>
            <Link href="/master/approve" 
                className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-700 transition duration-300"
            >
                Go to Approve User
            </Link>
        </div>
    );
};

export default MasterPage;