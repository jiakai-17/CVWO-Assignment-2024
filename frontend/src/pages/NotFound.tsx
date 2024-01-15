import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import { Link } from "react-router-dom";

export default function NotFound() {
  return (
    <>
      <Typography
        sx={{
          textAlign: "center",
          mt: 8,
        }}
        variant={"h4"}
      >
        404 Not Found
      </Typography>
      <div className={"mt-10 flex justify-center"}>
        <Link to={"/"}>
          <Button
            variant="outlined"
            disableElevation
            size="large"
          >
            Back to Home
          </Button>
        </Link>
      </div>
    </>
  );
}
