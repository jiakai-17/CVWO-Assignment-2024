"use client";

import * as React from "react";
import AuthPage from "@/components/AuthPage";

export default function Page() {
  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    console.log({
      email: data.get("email"),
      password: data.get("password"),
    });
  };

  return (
    <AuthPage handleSubmit={handleSubmit} type="signup" />
  );
}
