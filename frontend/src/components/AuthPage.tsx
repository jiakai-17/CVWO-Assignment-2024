"use client";

import * as React from "react";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Box from "@mui/material/Box";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import Link from "next/link";

export default function AuthPage(props: Readonly<{
  handleSubmit: (event: React.FormEvent<HTMLFormElement>) => void
  type: "login" | "signup"
}>) {

  const title = props.type === "login" ? "Login" : "Sign Up";
  const mainButtonLabel = props.type === "login" ? "Login" : "Create account";
  const secondaryButtonLink = props.type === "login" ? "/signup" : "/login";
  const secondaryButtonLabel = props.type === "login" ? "No account? Sign up" : "Already have an account? Login";

  return (
    <Container component="main" maxWidth="xs">
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Avatar sx={{ m: 1, bgcolor: "#f50057" }}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h4" sx={{ my: 2 }}>
          {title}
        </Typography>
        <Box component="form" onSubmit={props.handleSubmit} noValidate sx={{ mt: 1 }}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="username"
            label="Username"
            name="username"
            autoComplete="username"
            autoFocus
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            size="large"
            sx={{ mt: 3, mb: 2 }}
          >
            {mainButtonLabel}
          </Button>
          <Link href={secondaryButtonLink}>
            <Button
              fullWidth
              variant="outlined"
              size="large"
              sx={{ my: 2 }}>
              {secondaryButtonLabel}
            </Button>
          </Link>
        </Box>
      </Box>
    </Container>
  );
}
