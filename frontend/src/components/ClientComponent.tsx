"use client";

import * as React from "react";
import Navbar from "@/components/Navbar";

export default function ClientComponent({ children }: Readonly<{ children: React.ReactNode }>) {
  return (
    <>
      <Navbar />
      {children}
    </>
  );
}
