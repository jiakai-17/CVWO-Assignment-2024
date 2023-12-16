import * as React from "react";
import { useEffect, useState } from "react";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import UserAvatarDetails from "./UserAvatarDetails";
import { Link } from "react-router-dom";

export default function Navbar() {
  const [isLogin, setIsLogin] = useState(JSON.parse(localStorage?.getItem("isLogin") ?? "false"));

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
    console.log("handleLoginToggle", event);
    toggleLogin();
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
              <Button color="inherit">
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
                creator={"guest"}
                textColor={"white"}
                fontSize={"1rem"}
              />
            </div>
          )}
          <Button
            color="inherit"
            onClick={handleLoginToggle}
          >
            <Typography className={"text-white"}>{isLogin ? "Logout" : "Login"}</Typography>
          </Button>
        </Toolbar>
      </AppBar>
    </div>
  );
}
