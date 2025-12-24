'use client';

import React, { useRef, useMemo, Suspense } from 'react';
import { Canvas, useFrame } from '@react-three/fiber';
import { Stars, Float, MeshDistortMaterial, Sphere, Preload } from '@react-three/drei';
import * as THREE from 'three';

const Planet = () => {
  const planetRef = useRef<THREE.Mesh>(null);
  const atmosphereRef = useRef<THREE.Mesh>(null);

  useFrame((state) => {
    if (planetRef.current) {
      planetRef.current.rotation.y += 0.0005;
    }
    if (atmosphereRef.current) {
      atmosphereRef.current.rotation.y += 0.0007;
    }
  });

  return (
    <group position={[4, -2, -1]}>
      {/* Main Planet Body */}
      <mesh ref={planetRef}>
        <sphereGeometry args={[5, 64, 64]} />
        <meshStandardMaterial 
          color="#0a1a2a" 
          emissive="#020510"
          roughness={0.9}
          metalness={0.1}
        />
      </mesh>

      {/* Atmosphere Glow */}
      <mesh ref={atmosphereRef} scale={1.02}>
        <sphereGeometry args={[5, 64, 64]} />
        <meshPhongMaterial
          color="#4ade80"
          transparent
          opacity={0.15}
          side={THREE.BackSide}
          blending={THREE.AdditiveBlending}
        />
      </mesh>
      
      {/* Clouds / Surface Detail Distort */}
      <Sphere args={[5.02, 64, 64]}>
        <MeshDistortMaterial
          color="#ffffff"
          opacity={0.05}
          transparent
          distort={0.2}
          speed={0.5}
        />
      </Sphere>
    </group>
  );
};

const Debris = ({ count = 20 }) => {
  const debrisItems = useMemo(() => {
    return Array.from({ length: count }).map(() => ({
      position: [
        (Math.random() - 0.5) * 20,
        (Math.random() - 0.5) * 15,
        (Math.random() - 0.5) * 15,
      ] as [number, number, number],
      rotation: [Math.random() * Math.PI, Math.random() * Math.PI, Math.random() * Math.PI] as [number, number, number],
      scale: Math.random() * 0.3 + 0.05,
      speed: Math.random() * 0.1 + 0.05,
      type: Math.floor(Math.random() * 3), // 0: box, 1: tetra, 2: torus
    }));
  }, [count]);

  return (
    <>
      {debrisItems.map((item, i) => (
        <Float key={i} speed={item.speed * 5} rotationIntensity={2} floatIntensity={2}>
          <mesh position={item.position} rotation={item.rotation} scale={item.scale}>
            {item.type === 0 && <boxGeometry args={[1, 1, 1]} />}
            {item.type === 1 && <tetrahedronGeometry args={[1, 0]} />}
            {item.type === 2 && <torusGeometry args={[0.5, 0.2, 8, 16]} />}
            <meshStandardMaterial 
              color="#333" 
              metalness={0.9} 
              roughness={0.1} 
              emissive="#111"
            />
          </mesh>
        </Float>
      ))}
    </>
  );
};

const SpaceBackground = () => {
  return (
    <div className="fixed inset-0 -z-10 bg-black overflow-hidden">
      <Suspense fallback={<div className="w-full h-full bg-black" />}>
        <Canvas camera={{ position: [0, 0, 8], fov: 50 }} dpr={[1, 2]}>
          <ambientLight intensity={0.3} />
          <pointLight position={[15, 5, 5]} intensity={2} color="#4ade80" />
          <pointLight position={[-10, -10, -10]} intensity={1} color="#60a5fa" />
          
          <Stars radius={100} depth={50} count={7000} factor={4} saturation={0} fade speed={0.5} />
          
          <Planet />
          <Debris count={40} />
          
          <fog attach="fog" args={['#000', 8, 20]} />
          <Preload all />
        </Canvas>
      </Suspense>
    </div>
  );
};

export default SpaceBackground;
