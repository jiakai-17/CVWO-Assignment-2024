"use client";

import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import { useEffect, useState } from "react";
import Link from "next/link";
import UserAvatarDetails from "@/components/UserAvatarDetails";

export default function Navbar() {

  const [isLogin, setIsLogin] =
    useState(JSON.parse(localStorage?.getItem("isLogin") ?? "false"));

  useEffect(() => {
    if (localStorage?.getItem("isLogin") === null) {
      localStorage.setItem("isLogin", JSON.stringify(false));
    }
  }, [isLogin]);

  function toggleLogin() {
    localStorage.setItem("isLogin", JSON.stringify(!isLogin));
    setIsLogin(!isLogin);
  }

  function handleLoginToggle(event: React.MouseEvent<HTMLButtonElement>) {
    console.log("handleLoginToggle");
    toggleLogin();
  }

  return (
    <Box sx={{ mb: 8 }}>
      <AppBar position="static">
        <Toolbar>
          <Box sx={{ px: 2 }}>
            <Link href={"/"} style={{ textDecoration: "none", color: "#000" }}>
              <Typography variant="h6" sx={{ color: "white" }}>
                FORUM
              </Typography>
            </Link>
          </Box>
          <div style={{ flexGrow: 1 }} />
          {!isLogin &&
            <Button color="inherit">
              Signup
            </Button>
          }
          {isLogin &&
            <Box sx={{ display: "flex", gap: 1, alignItems: "center", mr: 2 }}>
              <Typography sx={{ color: "white" }}>
                Welcome, </Typography>
              <UserAvatarDetails creator={"guest"} textColor={"white"} fontSize={"1rem"} />
            </Box>
          }
          <Button color="inherit" onClick={handleLoginToggle} sx={{ ml: 2 }}>
            {isLogin ? "Logout" : "Login"}
          </Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
