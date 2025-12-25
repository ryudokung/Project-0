'use client';

import Image from "next/image";
import { useAuthSync } from "@/hooks/use-auth-sync";
import { useState } from "react";

interface LandingStageProps {
  onStart: () => void;
}

export default function LandingStage({ onStart }: LandingStageProps) {
  const { user, guestLogin, traditionalLogin, signup, isLoading } = useAuthSync();
  const [showAuthForm, setShowAuthForm] = useState<'LOGIN' | 'SIGNUP' | null>(null);
  const [showOptions, setShowOptions] = useState(false);
  const [formData, setFormData] = useState({ username: '', email: '', password: '' });

  const handleAuthSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (showAuthForm === 'LOGIN') {
      const res = await traditionalLogin(formData.username, formData.password);
      if (res.success) onStart();
    } else {
      const res = await signup(formData.username, formData.email, formData.password);
      if (res.success) onStart();
    }
  };

  return (
    <div className="flex min-h-screen flex-col items-center font-sans text-white bg-black">
      <header className="w-full max-w-7xl flex justify-between items-center p-6 z-20">
        <div className="flex items-center gap-2">
          <div className="relative h-8 w-8">
            <Image className="dark:invert" src="/next.svg" alt="Logo" fill />
          </div>
          <span className="font-bold tracking-tighter text-xl">PROJECT-0</span>
        </div>
        {user && (
          <div className="flex flex-col items-end">
            <span className="text-[10px] text-zinc-500 font-mono uppercase tracking-widest">Operative: {user.username}</span>
            <span className="text-[8px] text-pink-500 font-mono uppercase tracking-widest">Status: Online</span>
          </div>
        )}
      </header>

      <main className="flex flex-col items-center justify-center w-full max-w-7xl px-6 flex-1 z-10">
        <div className="flex flex-col items-center gap-8 text-center mb-20">
          <h1 className="text-5xl font-bold tracking-tighter sm:text-7xl uppercase italic">
            PROJECT-<span className="text-pink-500">0</span>
          </h1>
          <p className="max-w-[600px] text-zinc-400 md:text-xl font-light tracking-wide">
            The next generation of AI-driven space exploration and combat. 
            Create your pilot, command your fleet, and conquer the void.
          </p>
          
          <div className="flex flex-col gap-4 items-center w-full max-w-xs">
            {user ? (
              <button 
                onClick={onStart}
                className="w-full px-8 py-4 bg-white text-black font-black italic uppercase tracking-tighter hover:bg-gray-200 transition-all transform hover:-translate-y-1 shadow-[0_0_20px_rgba(255,255,255,0.2)]"
              >
                Enter the Void
              </button>
            ) : !showOptions ? (
              <button 
                onClick={() => setShowOptions(true)}
                className="w-full px-8 py-4 bg-white text-black font-black italic uppercase tracking-tighter hover:bg-gray-200 transition-all transform hover:-translate-y-1 shadow-[0_0_20px_rgba(255,255,255,0.2)]"
              >
                Initialize Neural Link
              </button>
            ) : !showAuthForm ? (
              <>
                <button 
                  onClick={guestLogin}
                  disabled={isLoading}
                  className="w-full px-8 py-4 bg-white text-black font-black italic uppercase tracking-tighter hover:bg-gray-200 transition-all transform hover:-translate-y-1 shadow-[0_0_20px_rgba(255,255,255,0.2)] disabled:opacity-50"
                >
                  {isLoading ? "Initializing..." : "Guest Access"}
                </button>
                
                <div className="flex items-center w-full gap-4">
                  <div className="h-[1px] bg-zinc-800 flex-1" />
                  <span className="text-[10px] text-zinc-600 font-bold uppercase tracking-widest">OR</span>
                  <div className="h-[1px] bg-zinc-800 flex-1" />
                </div>

                <button 
                  onClick={() => setShowAuthForm('LOGIN')}
                  className="w-full px-8 py-4 bg-zinc-900 hover:bg-zinc-800 text-white font-bold uppercase tracking-widest transition-all border border-zinc-800"
                >
                  Operative Login
                </button>
                
                <button 
                  className="w-full px-8 py-4 bg-pink-600/20 hover:bg-pink-600/30 text-pink-500 font-bold uppercase tracking-widest transition-all border border-pink-500/30 opacity-50 cursor-not-allowed"
                  title="Social Login Coming Soon"
                >
                  Social Login
                </button>
                
                <button 
                  onClick={() => setShowOptions(false)}
                  className="text-[10px] text-zinc-500 hover:text-white uppercase tracking-widest mt-2"
                >
                  Back
                </button>
              </>
            ) : (
              <form onSubmit={handleAuthSubmit} className="w-full flex flex-col gap-4 bg-zinc-900/50 p-6 border border-zinc-800 backdrop-blur-md">
                <h2 className="text-xl font-black italic uppercase tracking-tighter text-pink-500 mb-2">
                  {showAuthForm === 'LOGIN' ? 'Neural Link Access' : 'New Operative Registration'}
                </h2>
                
                <div className="space-y-2">
                  <input
                    type="text"
                    placeholder="USERNAME"
                    className="w-full bg-black border border-zinc-800 p-3 text-xs font-mono text-white focus:border-pink-500 outline-none"
                    value={formData.username}
                    onChange={(e) => setFormData({ ...formData, username: e.target.value })}
                    required
                  />
                  {showAuthForm === 'SIGNUP' && (
                    <input
                      type="email"
                      placeholder="EMAIL"
                      className="w-full bg-black border border-zinc-800 p-3 text-xs font-mono text-white focus:border-pink-500 outline-none"
                      value={formData.email}
                      onChange={(e) => setFormData({ ...formData, email: e.target.value })}
                      required
                    />
                  )}
                  <input
                    type="password"
                    placeholder="PASSWORD"
                    className="w-full bg-black border border-zinc-800 p-3 text-xs font-mono text-white focus:border-pink-500 outline-none"
                    value={formData.password}
                    onChange={(e) => setFormData({ ...formData, password: e.target.value })}
                    required
                  />
                </div>

                <button
                  type="submit"
                  disabled={isLoading}
                  className="w-full bg-white text-black py-3 font-black uppercase tracking-tighter hover:bg-pink-500 hover:text-white transition-all"
                >
                  {isLoading ? 'Processing...' : (showAuthForm === 'LOGIN' ? 'Authorize' : 'Register')}
                </button>

                <div className="flex justify-between items-center mt-2">
                  <button
                    type="button"
                    onClick={() => setShowAuthForm(showAuthForm === 'LOGIN' ? 'SIGNUP' : 'LOGIN')}
                    className="text-[10px] text-zinc-500 hover:text-white uppercase tracking-widest"
                  >
                    {showAuthForm === 'LOGIN' ? 'Need an account?' : 'Already registered?'}
                  </button>
                  <button
                    type="button"
                    onClick={() => setShowAuthForm(null)}
                    className="text-[10px] text-red-500 hover:text-red-400 uppercase tracking-widest"
                  >
                    Cancel
                  </button>
                </div>
              </form>
            )}
          </div>
        </div>

        {/* Decorative Elements */}
        <div className="absolute bottom-0 left-0 w-full h-64 bg-gradient-to-t from-pink-500/10 to-transparent pointer-events-none" />
      </main>
      
      <footer className="py-10 text-sm text-zinc-600 z-10">
        <p>© 2024 Project-0. All rights reserved.</p>
      </footer>
    </div>
  );
}
