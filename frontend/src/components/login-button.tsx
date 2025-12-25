'use client';

import { useAuthSync } from '@/hooks/use-auth-sync';
import { useState } from 'react';

export function LoginButton() {
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
        <div className="flex flex-col gap-2">
          <button
            onClick={() => setShowLocalForm('LOGIN')}
            className="text-[10px] text-zinc-400 hover:text-white font-bold uppercase tracking-widest text-left"
          >
            Login
          </button>
          <button
            onClick={() => setShowLocalForm('SIGNUP')}
            className="text-[10px] text-zinc-400 hover:text-white font-bold uppercase tracking-widest text-left"
          >
            Signup
          </button>
        </div>
      ) : (
        <form onSubmit={handleLocalSubmit} className="flex flex-col gap-2 bg-zinc-900/50 p-4 border border-zinc-800">
          <input
            type="text"
            placeholder="Username"
            className="bg-black border border-zinc-800 text-[10px] p-2 focus:outline-none focus:border-pink-500"
            value={formData.username}
            onChange={(e) => setFormData({ ...formData, username: e.target.value })}
            required
          />
          {showLocalForm === 'SIGNUP' && (
            <input
              type="email"
              placeholder="Email"
              className="bg-black border border-zinc-800 text-[10px] p-2 focus:outline-none focus:border-pink-500"
              value={formData.email}
              onChange={(e) => setFormData({ ...formData, email: e.target.value })}
              required
            />
          )}
          <input
            type="password"
            placeholder="Password"
            className="bg-black border border-zinc-800 text-[10px] p-2 focus:outline-none focus:border-pink-500"
            value={formData.password}
            onChange={(e) => setFormData({ ...formData, password: e.target.value })}
            required
          />
          <div className="flex gap-2">
            <button type="submit" className="flex-1 bg-white text-black text-[10px] font-bold py-2 uppercase">
              {showLocalForm}
            </button>
            <button type="button" onClick={() => setShowLocalForm(null)} className="px-4 border border-zinc-800 text-[10px] uppercase">
              Cancel
            </button>
          </div>
        </form>
      )}
    </div>
  );
}
