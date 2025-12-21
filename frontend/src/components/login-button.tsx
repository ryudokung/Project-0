'use client';

import {usePrivy} from '@privy-io/react-auth';

export function LoginButton() {
  const {login, logout, authenticated, user} = usePrivy();

  if (authenticated) {
    return (
      <div className="flex flex-col items-center gap-4">
        <p className="text-sm text-gray-400">
          Logged in as: <span className="text-white font-mono">{user?.wallet?.address || user?.email?.address}</span>
        </p>
        <button
          onClick={logout}
          className="px-6 py-2 bg-red-600 hover:bg-red-700 text-white rounded-lg transition-colors"
        >
          Logout
        </button>
      </div>
    );
  }

  return (
    <button
      onClick={login}
      className="px-8 py-3 bg-indigo-600 hover:bg-indigo-700 text-white font-bold rounded-xl shadow-lg shadow-indigo-500/20 transition-all transform hover:scale-105"
    >
      Enter Project-0
    </button>
  );
}
