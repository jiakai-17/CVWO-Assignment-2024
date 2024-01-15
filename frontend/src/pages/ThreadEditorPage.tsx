import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { Alert, Autocomplete, Chip, CircularProgress, Divider } from "@mui/material";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import PublishIcon from "@mui/icons-material/Publish";
import { Link, useNavigate } from "react-router-dom";
import Thread from "../models/Thread.tsx";
import { useContext, useEffect, useState } from "react";
import AuthContext from "../contexts/AuthContext.tsx";

export default function ThreadEditorPage(
  props: Readonly<{
    type: "create" | "edit";
  }>,
) {
  const [threadToEdit, setThreadToEdit] = useState<Thread | null>(null);
  const [isAuthorized, setIsAuthorized] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const { auth, isLoaded } = useContext(AuthContext);
  const navigate = useNavigate();

  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");
  const [tags, setTags] = useState<string[]>([] as string[]);

  const [isError, setIsError] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  const checkInvalidTitleAndBody = () => {
    const trimmedTitle = title.trim();
    const trimmedBody = body.trim();
    if (trimmedTitle.length < 1 || trimmedTitle.length > 100) {
      setIsError(true);
      setErrorMessage("Title must be between 1 and 100 characters.");
      return true;
    }
    if (trimmedBody.length < 1 || trimmedBody.length > 3000) {
      setIsError(true);
      setErrorMessage("Body must be between 1 and 3000 characters.");
      return true;
    }
    setIsError(false);
    setErrorMessage("");
    return false;
  };

  useEffect(() => {
    setIsLoading(true);
    if (!isLoaded) {
      return;
    }

    if (!auth.isLogin) {
      console.error("Not logged in");
      navigate("/login");
      return;
    }

    if (props.type === "create") {
      setIsLoading(false);
      setIsAuthorized(true);
      setThreadToEdit(null);
    }

    if (props.type === "edit") {
      setIsLoading(true);
      const id = window.location.pathname.split("/")[2];
      fetch(`/api/v1/thread/${id}`).then((res) => {
        if (!res.ok) {
          res.text().then((text) => {
            setIsError(true);
            setErrorMessage(text);
            setIsLoading(false);
          });
          setThreadToEdit(null);
        } else {
          res.json().then((data) => {
            if (data.creator !== auth.username) {
              navigate("/login");
            }
            setThreadToEdit(data);
            setTitle(data.title);
            setBody(data.body);
            setTags(data.tags);
            setIsAuthorized(true);
            setIsLoading(false);
          });
        }
      });
    }
  }, [auth, isLoaded, navigate, props.type]);

  const handleCreateThread = () => {
    setIsError(false);
    setErrorMessage("");
    if (checkInvalidTitleAndBody()) {
      return;
    }
    fetch("/api/v1/thread/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + auth.token,
      },
      body: JSON.stringify({
        title: title,
        body: body,
        tags: tags,
      }),
    }).then((res) => {
      if (!res.ok) {
        res.text().then((text) => {
          setIsError(true);
          setErrorMessage(text);
        });
      } else {
        res.json().then((data) => {
          navigate("/viewthread/" + data.id);
        });
      }
    });
  };

  const handleEditThread = () => {
    setIsError(false);
    setErrorMessage("");
    if (checkInvalidTitleAndBody()) {
      return;
    }
    const id = window.location.pathname.split("/")[2];
    fetch(`/api/v1/thread/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + auth.token,
      },
      body: JSON.stringify({
        title: title,
        body: body,
        tags: tags,
      }),
    }).then((res) => {
      if (!res.ok) {
        res.text().then((text) => {
          setIsError(true);
          setErrorMessage(text);
        });
      } else {
        navigate("/viewthread/" + id);
      }
    });
  };

  return (
    <>
      {isLoading && (
        <div className={"mt-16 flex flex-row justify-center"}>
          <CircularProgress />
          <Typography
            variant="h6"
            className={"px-6"}
          >
            Loading thread...
          </Typography>
        </div>
      )}

      {!isLoading && (
        <Box className={"mx-3 mb-10 mt-16 text-center"}>
          <Typography
            variant="h4"
            sx={{ px: 2 }}
          >
            {props.type === "edit" ? "Edit Thread" : "Create a new Thread"}
          </Typography>
          <Divider sx={{ m: 5 }} />
          {isError && (
            <Alert
              severity="error"
              sx={{ mx: 4, mb: 4 }}
            >
              {errorMessage}
            </Alert>
          )}
          {isAuthorized && (
            <Box
              sx={{
                mx: { xs: 2, sm: 4 },
                display: "flex",
                flexWrap: "wrap",
                flexDirection: "column",
              }}
            >
              <TextField
                id="outlined-basic"
                label="Title (max 100 chars)"
                variant="outlined"
                sx={{ my: 2, flexGrow: 1 }}
                defaultValue={props.type === "edit" ? threadToEdit?.title : ""}
                onChange={(event) => setTitle(event.target.value)}
                required={true}
              />
              <TextField
                id="outlined-multiline-static"
                label="Body (max 3000 chars)"
                multiline
                rows={10}
                variant="outlined"
                sx={{ my: 2, flexGrow: 1 }}
                defaultValue={props.type === "edit" ? threadToEdit?.body : ""}
                onChange={(event) => setBody(event.target.value)}
                required={true}
              />
              <Autocomplete
                multiple
                id="tags-outlined"
                defaultValue={props.type === "edit" ? threadToEdit?.tags : []}
                freeSolo
                renderTags={(value: readonly string[], getTagProps) =>
                  value.map((option: string, index: number) => (
                    <Chip
                      variant="outlined"
                      label={option}
                      {...getTagProps({ index })}
                    />
                  ))
                }
                sx={{ my: 2, flexGrow: 1, mb: 6 }}
                renderInput={(params) => (
                  <TextField
                    {...params}
                    variant="outlined"
                    label="Tags (optional)"
                    placeholder="Add up to 3 tags. Press Enter to create a tag."
                  />
                )}
                options={[]}
                onChange={(_event, value) => {
                  if (value.length > 3) {
                    value = value.slice(0, 3);
                  }
                  for (let i = 0; i < value.length; i++) {
                    value[i] = value[i].replace(/\s+/g, "-").toLowerCase();
                  }
                  setTags(value);
                }}
                value={tags}
              />

              <Box sx={{ display: "flex", justifyContent: "space-between" }}>
                <Link to={props.type === "edit" ? "/viewthread/" + threadToEdit?.id : "/"}>
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
                  onClick={props.type === "edit" ? handleEditThread : handleCreateThread}
                >
                  {props.type === "edit" ? "Save" : "Post"}
                </Button>
              </Box>
            </Box>
          )}
        </Box>
      )}
    </>
  );
}
