import { useContext, useEffect, useState } from "react";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import UserAvatarDetails from "./UserAvatarDetails";
import { Link, useNavigate } from "react-router-dom";
import AuthContext from "../contexts/AuthContext.tsx";

// Creates a navbar component to be displayed at the top of the page
export default function Navbar() {
  const navigate = useNavigate();
  const { auth, resetAuth, isLoaded } = useContext(AuthContext);
  const [isLogin, setIsLogin] = useState(auth.isLogin);

  useEffect(() => {
    if (!isLoaded) {
      return;
    }
    setIsLogin(auth.isLogin);
  }, [auth.isLogin, isLoaded]);

  const handleLogout = () => {
    resetAuth();
  };

  function handleLogin() {
    navigate("/login");
  }

  function handleSignup() {
    navigate("/signup");
  }

  return (
    <div>
      <AppBar
        position="static"
        className={"px-4"}
      >
        <Toolbar>
          <div className={"mr-8"}>
            <Link
              to={"/"}
              className={"no-underline"}
            >
              <Typography
                variant="h6"
                className={"text-white"}
              >
                FORUM
              </Typography>
            </Link>
          </div>
          <div className={"flex-grow"} />
          {!isLogin && (
            <div className={"mr-4"}>
              <Button
                color="inherit"
                onClick={handleSignup}
              >
                <Typography className={"text-white"}>Sign Up</Typography>
              </Button>
            </div>
          )}
          {isLogin && (
            <div className={"mr-6 flex min-w-[30px] items-center gap-2"}>
              <div className={"hidden  md:block"}>
                <Typography className={"text-white"}>Welcome,</Typography>
              </div>
              <UserAvatarDetails
                creator={auth.username}
                textColor={"white"}
                fontSize={"1rem"}
              />
            </div>
          )}
          <Button
            color="inherit"
            onClick={isLogin ? handleLogout : handleLogin}
          >
            <Typography className={"text-white"}>{isLogin ? "Logout" : "Login"}</Typography>
          </Button>
        </Toolbar>
      </AppBar>
    </div>
  );
}
