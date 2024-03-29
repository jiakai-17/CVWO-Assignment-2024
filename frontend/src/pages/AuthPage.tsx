import * as React from "react";
import { useContext, useState } from "react";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import { useNavigate } from "react-router-dom";
import { Alert, CircularProgress, Divider } from "@mui/material";
import AuthContext from "../contexts/AuthContext.tsx";

export default function AuthPage(
  props: Readonly<{
    type: "login" | "signup";
  }>,
) {
  const navigate = useNavigate();
  const { setAuthFromToken } = useContext(AuthContext);

  const title = props.type === "login" ? "Login" : "Sign Up";
  const mainButtonLabel = props.type === "login" ? "Login" : "Create account";
  const secondaryButtonLink = props.type === "login" ? "/signup" : "/login";
  const secondaryButtonLabel = props.type === "login" ? "No account? Sign up" : "Already have an account? Login";

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleUsernameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUsername(event.target.value);
  };

  const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value);
  };

  const [isInvalidUsername, setIsInvalidUsername] = useState(false);
  const [isInvalidPassword, setIsInvalidPassword] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [isError, setIsError] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  const checkInvalidUsername = () => {
    if (username.length < 1 || username.length > 30 || username.includes(" ")) {
      setIsInvalidUsername(true);
      return true;
    } else {
      setIsInvalidUsername(false);
      return false;
    }
  };

  const checkInvalidPassword = () => {
    if (password.length < 6) {
      setIsInvalidPassword(true);
      return true;
    } else {
      setIsInvalidPassword(false);
      return false;
    }
  };

  const handleSubmit = () => {
    setIsError(false);
    setErrorMessage("");
    setIsLoading(true);

    const isInvalidUsername = checkInvalidUsername();
    const isInvalidPassword = checkInvalidPassword();

    if (isInvalidUsername || isInvalidPassword) {
      setIsLoading(false);
      return;
    }

    const handleApiResponse = (response: Response) => {
      if (response.status === 200) {
        response.json().then((data) => {
          localStorage.setItem("token", data.token);
          setAuthFromToken(data.token);
          navigate("/");
        });
      } else {
        setIsLoading(false);
        response.text().then((text) => {
          setIsError(true);
          setErrorMessage(text);
        });
      }
    };

    if (props.type === "login") {
      fetch("/api/v1/user/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username,
          password,
        }),
      }).then(handleApiResponse);
    } else if (props.type === "signup") {
      fetch("/api/v1/user/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username,
          password,
        }),
      }).then(handleApiResponse);
    } else {
      throw new Error("Invalid auth page type");
    }
  };

  return (
    <div className={"mx-2 mb-10 mt-16 text-center"}>
      <Typography
        variant="h4"
        className={"px-2"}
      >
        {title}
      </Typography>
      <Divider sx={{ mx: 3, my: 6 }} />
      <div className={"mx-8 flex justify-items-center"}>
        <div className={"mx-auto mt-1 max-w-xl"}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="username"
            label="Username"
            name="username"
            autoComplete="username"
            autoFocus
            value={username}
            onChange={handleUsernameChange}
            error={isInvalidUsername}
            helperText={"Username must be between 1 and 30 characters long and cannot contain spaces"}
            onKeyDown={(event) => {
              if (event.key === "Enter") {
                handleSubmit();
              }
            }}
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
            value={password}
            onChange={handlePasswordChange}
            error={isInvalidPassword}
            helperText={"Password must be at least 6 characters long"}
            onKeyDown={(event) => {
              if (event.key === "Enter") {
                handleSubmit();
              }
            }}
          />
          {isError && (
            <Alert
              severity="error"
              className={"mt-3"}
            >
              {errorMessage}
            </Alert>
          )}
          <Button
            type="submit"
            fullWidth
            variant="contained"
            size="large"
            className={"h-11"}
            sx={{ mt: 3, mb: 2 }}
            onClick={handleSubmit}
            disabled={isLoading}
          >
            {isLoading ? <CircularProgress size={28} /> : mainButtonLabel}
          </Button>
          <Button
            fullWidth
            variant="outlined"
            size="large"
            sx={{ my: 2 }}
            className={"h-11"}
            onClick={() => {
              setUsername("");
              setPassword("");
              setIsInvalidPassword(false);
              setIsInvalidUsername(false);
              setIsError(false);
              setErrorMessage("");
              navigate(secondaryButtonLink);
            }}
            disabled={isLoading}
          >
            {secondaryButtonLabel}
          </Button>
        </div>
      </div>
    </div>
  );
}
