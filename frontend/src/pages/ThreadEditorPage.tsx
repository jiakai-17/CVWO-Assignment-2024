import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { Divider } from "@mui/material";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import PublishIcon from "@mui/icons-material/Publish";
import { Link, useLoaderData } from "react-router-dom";
import Thread from "../models/thread/Thread.tsx";

export default function ThreadEditorPage(
  props: Readonly<{
    type: "create" | "edit";
  }>,
) {
  const threadToEdit = useLoaderData() as Thread;

  return (
    <Box className={"mx-3 mb-10 mt-16 text-center"}>
      <Typography
        variant="h4"
        sx={{ px: 2 }}
      >
        {props.type === "edit" ? "Edit Thread" : "Create a new Thread"}
      </Typography>
      <Divider sx={{ m: 5 }} />
      <Box sx={{ mx: { xs: 2, sm: 4 }, display: "flex", flexWrap: "wrap", flexDirection: "column" }}>
        <TextField
          id="outlined-basic"
          label="Title"
          variant="outlined"
          sx={{ my: 2, flexGrow: 1 }}
          defaultValue={props.type === "edit" ? threadToEdit?.title : ""}
        />
        <TextField
          id="outlined-multiline-static"
          label="Body"
          multiline
          rows={10}
          variant="outlined"
          sx={{ my: 2, flexGrow: 1 }}
          defaultValue={props.type === "edit" ? threadToEdit?.body : ""}
        />
        <Box sx={{ display: "flex", justifyContent: "space-between" }}>
          <Link to={props.type === "edit" ? "/viewthread/" + threadToEdit.id : "/"}>
            <Button
              variant="outlined"
              size="large"
              color="error"
              sx={{ mr: 2 }}
            >
              Cancel
            </Button>
          </Link>
          <Button
            variant="contained"
            disableElevation
            size="large"
            endIcon={<PublishIcon />}
            sx={{ width: "12rem" }}
          >
            {props.type === "edit" ? "Save" : "Post"}
          </Button>
        </Box>
      </Box>
    </Box>
  );
}
