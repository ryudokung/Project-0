'use client';

import { usePrivy } from '@privy-io/react-auth';
import { useAuthSync } from '@/hooks/use-auth-sync';
import { useState } from 'react';

export function LoginButton() {
  const { login: privyLogin, authenticated: privyAuthenticated } = usePrivy();
  const { user, logout, guestLogin, traditionalLogin, signup, isLoading } = useAuthSync();
  const [showLocalForm, setShowLocalForm] = useState<'LOGIN' | 'SIGNUP' | null>(null);
  const [formData, setFormData] = useState({ username: '', email: '', password: '' });

  const handleLocalSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (showLocalForm === 'LOGIN') {
      await traditionalLogin(formData.username, formData.password);
    } else {
      await signup(formData.username, formData.email, formData.password);
    }
    setShowLocalForm(null);
  };

  if (user) {
    return (
      <div className="flex flex-col items-start gap-2">
        <p className="text-[10px] text-gray-500 font-mono uppercase tracking-widest">
          [OPERATIVE: {user.username}] // [TYPE: {user.auth_type}]
        </p>
        <button
          onClick={logout}
          className="text-[10px] text-red-500 hover:text-red-400 font-bold uppercase tracking-widest underline underline-offset-4"
        >
          Terminate Session
        </button>
      </div>
    );
  }

  return (
    <div className="flex flex-col gap-4">
      {!showLocalForm ? (
        <>
          <button
            onClick={guestLogin}
            disabled={isLoading}
            className="px-8 py-3 bg-white text-black font-black italic uppercase tracking-tighter hover:bg-gray-200 transition-all transform hover:-translate-y-1"
          >
            {isLoading ? 'Initializing...' : 'Quick Start (Guest)'}
          </button>
          
          <button
            onClick={privyLogin}
            className="px-8 py-3 bg-indigo-600 text-white font-bold uppercase tracking-widest text-xs hover:bg-indigo-500 transition-all"
          >
            Social Login (Google/Email)
          </button>

          <button
            onClick={() => setShowLocalForm('LOGIN')}
            className="text-[10px] text-gray-500 hover:text-white font-mono uppercase tracking-widest text-center"
          >
            Traditional Login
          </button>
        </>
      ) : (
        <form onSubmit={handleLocalSubmit} className="flex flex-col gap-2 bg-zinc-900 p-4 border border-zinc-800">
          <h3 className="text-xs font-bold uppercase tracking-widest text-indigo-400 mb-2">
            {showLocalForm === 'LOGIN' ? 'System Access' : 'New Operative'}
          </h3>
          <input
            type="text"
            placeholder="USERNAME"
            className="bg-black border border-zinc-800 p-2 text-xs font-mono text-white focus:border-indigo-500 outline-none"
            value={formData.username}
            onChange={(e) => setFormData({ ...formData, username: e.target.value })}
            required
          />
          {showLocalForm === 'SIGNUP' && (
            <input
              type="email"
              placeholder="EMAIL"
              className="bg-black border border-zinc-800 p-2 text-xs font-mono text-white focus:border-indigo-500 outline-none"
              value={formData.email}
              onChange={(e) => setFormData({ ...formData, email: e.target.value })}
              required
            />
          )}
          <input
            type="password"
            placeholder="PASSWORD"
            className="bg-black border border-zinc-800 p-2 text-xs font-mono text-white focus:border-indigo-500 outline-none"
            value={formData.password}
            onChange={(e) => setFormData({ ...formData, password: e.target.value })}
            required
          />
          <button
            type="submit"
            className="bg-indigo-600 text-white py-2 text-xs font-bold uppercase tracking-widest mt-2"
          >
            {showLocalForm === 'LOGIN' ? 'Authorize' : 'Register'}
          </button>
          <button
            type="button"
            onClick={() => setShowLocalForm(showLocalForm === 'LOGIN' ? 'SIGNUP' : 'LOGIN')}
            className="text-[9px] text-gray-500 hover:text-white uppercase tracking-widest mt-1"
          >
            {showLocalForm === 'LOGIN' ? 'Need an account?' : 'Already registered?'}
          </button>
          <button
            type="button"
            onClick={() => setShowLocalForm(null)}
            className="text-[9px] text-red-500 hover:text-red-400 uppercase tracking-widest mt-1"
          >
            Cancel
          </button>
        </form>
      )}
    </div>
  );
}
