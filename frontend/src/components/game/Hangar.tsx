'use client';

import ShowcaseEngine from '@/components/game/ShowcaseEngine';
import { useAuthSync } from '@/hooks/use-auth-sync';
import { useEffect, useState } from 'react';
import { LoginButton } from '@/components/login-button';
import { hangarSystem, HangarState } from '@/systems/HangarSystem';
import { gameEvents, GAME_EVENTS } from '@/systems/EventBus';
import { vehicleService, Item } from '@/services/vehicle';

interface HangarProps {
  onDeploy: () => void;
  onGacha: () => void;
}

export default function Hangar({ onDeploy, onGacha }: HangarProps) {
  const { user: backendUser } = useAuthSync();
  const [hangarState, setHangarState] = useState<HangarState>(hangarSystem.getState());
  const [items, setItems] = useState<Item[]>([]);
  const [repairCost, setRepairCost] = useState<{ [key: string]: number }>({});

  useEffect(() => {
    const unsubscribe = gameEvents.on(GAME_EVENTS.HANGAR_UPDATED, (newState: HangarState) => {
      setHangarState(newState);
    });

    if (backendUser?.active_character_id || backendUser?.id) {
      hangarSystem.refreshVehicles(backendUser.active_character_id || backendUser.id);
      loadItems();
    }

    return () => unsubscribe();
  }, [backendUser?.id, backendUser?.active_character_id]);

  const loadItems = async () => {
    try {
      const fetchedItems = await vehicleService.getItems();
      setItems(fetchedItems);
    } catch (error) {
      console.error("Failed to load items", error);
    }
  };

  const handleRepair = async (itemId: string) => {
    try {
      const result = await vehicleService.repairItem(itemId);
      alert(`Repaired for ${result.cost} Scrap!`);
      loadItems(); // Refresh
    } catch (error: any) {
      alert(`Repair failed: ${error.message}`);
    }
  };

  const handleMint = async (itemId: string) => {
    try {
      const result = await vehicleService.mintItem(itemId);
      alert(`Minted successfully! Token ID: ${result.token_id}`);
      loadItems(); // Refresh
    } catch (error: any) {
      alert(`Minting failed: ${error.message}`);
    }
  };

  const handleEquip = async (itemId: string) => {
    if (!currentVehicle) return;
    try {
      await vehicleService.equipItem(itemId, currentVehicle.id);
      loadItems();
    } catch (error: any) {
      alert(`Equip failed: ${error.message}`);
    }
  };

  const handleUnequip = async (itemId: string) => {
    try {
      await vehicleService.unequipItem(itemId);
      loadItems();
    } catch (error: any) {
      alert(`Unequip failed: ${error.message}`);
    }
  };

  const vehicles = hangarState.vehicles || [];
  const currentVehicle = vehicles.find(v => v.id === hangarState.selectedVehicleId) || vehicles[0];

  // Filter items equipped on current vehicle (assuming parent_item_id matches vehicle id)
  const equippedItems = items.filter(i => i.is_equipped && i.parent_item_id === currentVehicle?.id);
  const inventoryItems = items.filter(i => !i.is_equipped);

  return (
    <div className="relative w-full h-screen bg-black">
      <ShowcaseEngine vehicleId={currentVehicle?.id || "MCH-001-ALPHA"} />
      
      {/* UI Overlay */}
      <div className="absolute top-8 left-8 z-10">
        <div className="flex items-start gap-4">
          {/* Pilot ID Badge */}
          <div className="bg-zinc-900/80 border border-zinc-700 p-4 rounded-lg backdrop-blur-md shadow-xl flex gap-4 items-center">
            <div className="w-16 h-16 bg-zinc-800 border border-zinc-700 rounded overflow-hidden flex items-center justify-center">
              {backendUser?.active_character?.gender ? (
                <img 
                  src={`/images/pilots/${backendUser.active_character.gender.toLowerCase()}.png`} 
                  alt="Pilot Avatar"
                  className="w-full h-full object-cover"
                />
              ) : (
                <div className="text-pink-500 text-2xl font-bold">
                  {backendUser?.active_character?.name?.[0] || 'P'}
                </div>
              )}
            </div>
            <div>
              <h2 className="text-xl font-bold text-white tracking-tighter uppercase italic">
                {backendUser?.active_character?.name || 'UNREGISTERED PILOT'}
              </h2>
              <div className="flex gap-2 mt-1">
                <span className="text-[10px] font-mono text-zinc-400 bg-zinc-800 px-1.5 py-0.5 rounded">
                  {backendUser?.active_character?.gender || 'UNKNOWN'}
                </span>
                <span className="text-[10px] font-mono text-pink-400 bg-pink-400/10 border border-pink-400/20 px-1.5 py-0.5 rounded">
                  RANK 1
                </span>
              </div>
            </div>
          </div>
        </div>
        
        <div className="mt-4 flex flex-col gap-2">
          <LoginButton />

          {backendUser?.wallet_address && (
            <div className="text-[10px] text-green-500 font-mono bg-green-500/10 border border-green-500/20 px-2 py-1 rounded">
              WALLET: {backendUser.wallet_address.slice(0, 6)}...{backendUser.wallet_address.slice(-4)}
            </div>
          )}
        </div>
        {currentVehicle && (
          <div className="mt-4 flex gap-2">
            {currentVehicle.is_void_touched && (
              <span key="void-touched" className="px-2 py-1 bg-purple-900/50 border border-purple-500 text-purple-300 text-[10px] font-bold uppercase tracking-widest">
                Void-Touched
              </span>
            )}
            <span key="tier" className="px-2 py-1 bg-zinc-900 border border-zinc-700 text-zinc-400 text-[10px] font-bold uppercase tracking-widest">
              Tier {currentVehicle.tier || 1}
            </span>
            <span key="rarity" className="px-2 py-1 bg-zinc-900 border border-zinc-700 text-zinc-400 text-[10px] font-bold uppercase tracking-widest">
              {currentVehicle.rarity}
            </span>
          </div>
        )}
      </div>

      <div className="absolute bottom-8 right-8 z-10 flex flex-col gap-4">
        {/* Maintenance Panel */}
        <div className="bg-zinc-900/90 border border-zinc-700 p-4 rounded-lg backdrop-blur-md w-80 max-h-96 overflow-y-auto">
          <h3 className="text-white font-bold mb-2 uppercase tracking-wider border-b border-zinc-700 pb-1">Maintenance</h3>
          
          {equippedItems.length === 0 && inventoryItems.length === 0 && (
             <div className="text-zinc-500 text-xs">No items found.</div>
          )}

          <div className="flex flex-col gap-2">
            {/* Equipped Items */}
            {equippedItems.map(item => (
              <div key={item.id} className="bg-zinc-800/50 p-2 rounded border border-zinc-700">
                <div className="flex justify-between items-center mb-1">
                  <span className="text-zinc-300 text-sm font-bold">{item.name}</span>
                  <span className={`text-[10px] px-1 rounded ${item.condition === 'PRISTINE' ? 'bg-green-900 text-green-300' : 'bg-red-900 text-red-300'}`}>
                    {item.condition}
                  </span>
                </div>
                <div className="flex justify-between items-center text-[10px] text-zinc-500 mb-2">
                   <span>Durability: {item.durability}/{item.max_durability}</span>
                   {item.damage_type && <span className="text-blue-400">{item.damage_type}</span>}
                </div>
                
                <div className="flex gap-2 mb-2">
                  <button 
                    onClick={() => handleRepair(item.id)}
                    disabled={item.durability >= item.max_durability}
                    className="flex-1 bg-blue-600 hover:bg-blue-500 disabled:bg-zinc-700 disabled:text-zinc-500 text-white text-xs py-1 rounded"
                  >
                    Repair
                  </button>
                  <button 
                    onClick={() => handleMint(item.id)}
                    disabled={item.is_nft || item.rarity === 'COMMON'} // Simple check
                    className="flex-1 bg-purple-600 hover:bg-purple-500 disabled:bg-zinc-700 disabled:text-zinc-500 text-white text-xs py-1 rounded"
                  >
                    {item.is_nft ? 'Minted' : 'Mint NFT'}
                  </button>
                </div>
                <button 
                  onClick={() => handleUnequip(item.id)}
                  className="w-full bg-zinc-700 hover:bg-zinc-600 text-zinc-300 text-xs py-1 rounded"
                >
                  Unequip
                </button>
              </div>
            ))}
             {/* Inventory Items (Simplified) */}
             {inventoryItems.length > 0 && (
                <div className="mt-2 pt-2 border-t border-zinc-700">
                  <h4 className="text-zinc-500 text-xs uppercase mb-2">Inventory</h4>
                  {inventoryItems.map(item => (
                    <div key={item.id} className="bg-zinc-800/30 p-2 rounded border border-zinc-700/50 mb-2 opacity-75">
                        <div className="flex justify-between items-center mb-1">
                            <span className="text-zinc-400 text-xs">{item.name}</span>
                            <span className="text-[10px] text-zinc-500">{item.slot}</span>
                        </div>
                        <div className="flex gap-1">
                            <button 
                                onClick={() => handleEquip(item.id)}
                                className="flex-1 text-[10px] bg-green-900/50 hover:bg-green-800 text-green-300 px-2 py-0.5 rounded border border-green-900"
                            >
                                Equip
                            </button>
                            <button 
                                onClick={() => handleRepair(item.id)}
                                className="flex-1 text-[10px] bg-zinc-700 hover:bg-zinc-600 text-zinc-300 px-2 py-0.5 rounded"
                            >
                                Repair
                            </button>
                        </div>
                    </div>
                  ))}
                </div>
             )}
          </div>
        </div>

        <button 
          onClick={onGacha}
          className="bg-yellow-400 text-black px-6 py-2 font-bold uppercase tracking-tighter hover:bg-yellow-300 transition-colors text-center"
        >
          Void Signals (Gacha)
        </button>
        <button className="bg-white text-black px-6 py-2 font-bold uppercase tracking-tighter hover:bg-gray-200 transition-colors">
          Share to Social
        </button>
        <button 
          onClick={onDeploy}
          className="border border-white text-white px-6 py-2 font-bold uppercase tracking-tighter hover:bg-white hover:text-black transition-colors text-center"
        >
          Launch Expedition
        </button>
      </div>
    </div>
  );
}
