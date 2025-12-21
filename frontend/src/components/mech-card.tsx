'use client';

import Image from 'next/image';

interface MechStats {
  hp: number;
  attack: number;
  defense: number;
  speed: number;
}

interface Mech {
  id: string;
  vehicle_type: string;
  class: string;
  image_url: string;
  stats: MechStats;
  rarity: string;
  season: string;
}

export function MechCard({ mech }: { mech: Mech }) {
  return (
    <div className="group relative overflow-hidden rounded-2xl bg-zinc-900 border border-zinc-800 p-4 transition-all hover:border-indigo-500/50 hover:shadow-2xl hover:shadow-indigo-500/10">
      <div className="relative aspect-square w-full overflow-hidden rounded-xl bg-zinc-800 mb-4">
        <Image
          src={mech.image_url}
          alt={`${mech.vehicle_type} ${mech.class}`}
          fill
          className="object-cover transition-transform group-hover:scale-110"
        />
        <div className="absolute top-2 right-2">
          <span className={`px-2 py-1 rounded text-[10px] font-bold uppercase tracking-wider ${
            mech.rarity === 'LEGENDARY' ? 'bg-yellow-500 text-black' :
            mech.rarity === 'RARE' ? 'bg-blue-500 text-white' :
            'bg-zinc-700 text-zinc-300'
          }`}>
            {mech.rarity}
          </span>
        </div>
      </div>

      <div className="space-y-3">
        <div>
          <h3 className="text-lg font-bold text-white leading-tight">
            {mech.vehicle_type} <span className="text-indigo-400">{mech.class}</span>
          </h3>
          <p className="text-xs text-zinc-500 uppercase tracking-widest">{mech.season} Edition</p>
        </div>

        <div className="grid grid-cols-2 gap-2">
          <Stat label="HP" value={mech.stats.hp} color="bg-green-500" />
          <Stat label="ATK" value={mech.stats.attack} color="bg-red-500" />
          <Stat label="DEF" value={mech.stats.defense} color="bg-blue-500" />
          <Stat label="SPD" value={mech.stats.speed} color="bg-yellow-500" />
        </div>
      </div>
    </div>
  );
}

function Stat({ label, value, color }: { label: string; value: number; color: string }) {
  return (
    <div className="bg-zinc-800/50 rounded-lg p-2 border border-zinc-700/50">
      <div className="flex justify-between items-center mb-1">
        <span className="text-[10px] text-zinc-500 font-bold">{label}</span>
        <span className="text-xs font-mono text-white">{value}</span>
      </div>
      <div className="h-1 w-full bg-zinc-700 rounded-full overflow-hidden">
        <div 
          className={`h-full ${color}`} 
          style={{ width: `${Math.min(value, 100)}%` }}
        />
      </div>
    </div>
  );
}
