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

  const [isLogin, setIsLogin] = useState(false);

  function toggleLogin() {
    setIsLogin(!isLogin);
  }

  function handleLoginToggle(event: React.MouseEvent<HTMLButtonElement>) {
    console.log("handleLoginToggle");
    toggleLogin();
  }

  return (
    <Box sx={{ flexGrow: 1, mb: 8 }}>
      <AppBar position="static">
        <Toolbar>
          <Box sx={{ flexGrow: 1, px: 2 }}>
            <Link href={"/"} style={{ textDecoration: "none", color: "#000" }}>
              <Typography variant="h6" sx={{ color: "white" }}>
                FORUM
              </Typography>
            </Link>
          </Box>
          <Button color="inherit" onClick={handleLoginToggle}>
            {isLogin ? "Logout" : "Login"}
          </Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
