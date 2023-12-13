"use client";

import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import { useState } from "react";
import Link from "next/link";

export default function Navbar() {

  const currentLoginState = localStorage.getItem("isLogin") ?? "false";

  const [isLogin, setIsLogin] = useState(JSON.parse(currentLoginState));
  localStorage.setItem("isLogin", currentLoginState);

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
          {!isLogin && <Button color="inherit">
            Signup
          </Button>}
          <Button color="inherit" onClick={handleLoginToggle} sx={{ ml: 2 }}>
            {isLogin ? "Logout" : "Login"}
          </Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
