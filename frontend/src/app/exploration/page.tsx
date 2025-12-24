'use client';

import { useState, useEffect } from 'react';

export default function ExplorationPage() {
  const [fuel, setFuel] = useState(100);
  const [sector, setSector] = useState('SOL-GATE');
  const [status, setStatus] = useState('IN_TRANSIT');
  const [logs, setLogs] = useState<string[]>(['Mothership engines engaged.', 'Scanning for anomalies...']);

  const handleJump = () => {
    if (fuel >= 10) {
      setFuel(prev => prev - 10);
      setSector('KAIROS-7');
      setLogs(prev => [...prev, 'Jump successful. Arrived at KAIROS-7.', 'Warning: High radiation detected.']);
    }
  };

  return (
    <main className="min-h-screen bg-black text-green-500 font-mono p-8 flex flex-col gap-8">
      <div className="border border-green-900 p-4 bg-green-950/10">
        <h1 className="text-2xl font-bold tracking-tighter uppercase mb-4">
          Mothership Command Console
        </h1>
        
        <div className="grid grid-cols-3 gap-4 mb-8">
          <div className="border border-green-800 p-2">
            <div className="text-xs text-green-700 uppercase">Current Sector</div>
            <div className="text-xl font-bold">{sector}</div>
          </div>
          <div className="border border-green-800 p-2">
            <div className="text-xs text-green-700 uppercase">Fuel Level</div>
            <div className="text-xl font-bold">{fuel}%</div>
            <div className="w-full bg-green-900 h-1 mt-1">
              <div className="bg-green-500 h-full" style={{ width: `${fuel}%` }}></div>
            </div>
          </div>
          <div className="border border-green-800 p-2">
            <div className="text-xs text-green-700 uppercase">Status</div>
            <div className="text-xl font-bold animate-pulse">{status}</div>
          </div>
        </div>

        <div className="flex gap-4">
          <button 
            onClick={handleJump}
            className="bg-green-500 text-black px-6 py-2 font-bold uppercase hover:bg-green-400 transition-colors"
          >
            Initiate Jump (10% Fuel)
          </button>
          <button className="border border-green-500 px-6 py-2 font-bold uppercase hover:bg-green-500 hover:text-black transition-colors">
            Deep Scan
          </button>
        </div>
      </div>

      <div className="flex-1 border border-green-900 p-4 bg-green-950/5 overflow-y-auto">
        <div className="text-xs text-green-700 uppercase mb-2">Mission Logs</div>
        {logs.map((log, i) => (
          <div key={`log-${i}`} className="text-sm mb-1">
            <span className="text-green-800 mr-2">[{new Date().toLocaleTimeString()}]</span>
            {log}
          </div>
        ))}
      </div>
    </main>
  );
}
