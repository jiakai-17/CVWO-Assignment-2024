import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { CircularProgress, Divider, ListItemText } from "@mui/material";
import { useContext, useEffect, useState } from "react";
import { Link } from "react-router-dom";
import ThreadTag from "../components/ThreadTag.tsx";
import UserContentTimestamp from "../components/UserContentTimestamp.tsx";
import UserAvatarDetails from "../components/UserAvatarDetails.tsx";
import SortButton from "../components/SortButton.tsx";
import sampleComments from "../models/comment/SampleComments.tsx";
import Thread from "../models/thread/Thread.tsx";
import CommentTextField from "../components/CommentTextField.tsx";
import { ThreadComment } from "./ThreadComment.tsx";
import authContext from "../contexts/AuthContext.tsx";

export default function ThreadPage() {
  const [isLoading, setIsLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState("");
  const [threadToDisplay, setThreadToDisplay] = useState<Thread | null>(null);
  const { auth } = useContext(authContext);

  useEffect(() => {
    const id = window.location.pathname.split("/")[2];
    fetch(`/api/v1/thread/${id}`).then((res) => {
      if (!res.ok) {
        setIsLoading(false);
        res.text().then((text) => setErrorMessage(text));
        setThreadToDisplay(null);
      } else {
        setIsLoading(false);
        setErrorMessage("");
        res.json().then((data) => setThreadToDisplay(data));
      }
    });
  }, []);

  const [commentContent, setCommentContent] = useState("");

  useEffect(() => {
    console.log("comment content changed", commentContent);
  }, [commentContent]);

  // Sorting comments by criteria
  const [commentSortCriteria, setCommentSortCriteria] = useState("Newest first");

  const availableCommentSortCriteria = new Map<string, string>([
    ["Newest first", "date desc"],
    ["Oldest first", "date asc"],
  ]);

  useEffect(() => {
    console.log(commentSortCriteria);
  }, [commentSortCriteria]);

  return (
    <Box className={"mt-16"}>
      <Box sx={{ mx: 4, mb: 4 }}>
        <Link to={"/"}>
          <Button
            variant="outlined"
            disableElevation
            size="large"
          >
            Back to Home
          </Button>
        </Link>
      </Box>

      {isLoading && (
        <div className={"flex flex-row justify-center"}>
          <CircularProgress />
          <Typography
            variant="h6"
            className={"px-6"}
          >
            Loading thread...
          </Typography>
        </div>
      )}

      {!isLoading && threadToDisplay === null && (
        <Box
          sx={{
            mx: 2,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Typography
            variant="h4"
            sx={{ mb: 6 }}
          >
            {errorMessage}
          </Typography>
        </Box>
      )}

      {!isLoading && threadToDisplay !== null && (
        <>
          <Box
            sx={{
              mx: 4,
              maxWidth: "full",
              borderRadius: "0.375rem",
              border: 2,
              borderColor: "lightgray",
              px: 3,
            }}
          >
            <Typography
              sx={{
                fontSize: { xs: "1.125rem", sm: "1.25rem" },
                fontWeight: "600",
                py: 2,
                overflowWrap: "break-word",
              }}
            >
              {threadToDisplay.title}
            </Typography>
            <Divider />
            <Typography
              sx={{
                fontSize: { xs: "0.875rem", sm: "1rem" },
                overflowWrap: "break-word",
                py: 2,
                mb: 5,
                whiteSpace: "pre-line",
              }}
            >
              {threadToDisplay.body}
            </Typography>
            <Divider />
            <Box
              sx={{
                display: "inline-flex",
                gap: "0.5rem",
                flexWrap: "wrap",
                alignItems: "center",
              }}
            >
              <Typography
                sx={{
                  fontSize: "1rem",
                  py: 2,
                  color: "rgba(0, 0, 0, 0.6)",
                }}
              >
                {threadToDisplay.tags.length === 0 ? "Tags: None" : "Tags:"}
              </Typography>
              {threadToDisplay.tags.map((tag) => (
                <ThreadTag
                  key={threadToDisplay.id + tag}
                  tag={tag}
                />
              ))}
            </Box>
            <Divider />
            <Box
              aria-label={"Thread Details"}
              sx={{
                display: "inline-flex",
                flexDirection: { xs: "column", sm: "row" },
                flexWrap: "wrap",
                width: "100%",
                my: 2,
              }}
            >
              <Box
                aria-label={"Thread Creator"}
                sx={{
                  display: "flex",
                  gap: { xs: "0.4rem", sm: "0.5rem" },
                  alignItems: "center",
                  maxWidth: "100%",
                  pr: 1,
                }}
              >
                <ListItemText
                  secondary={"Posted by "}
                  sx={{
                    fontSize: { xs: "0.5rem", sm: "body2" },
                    flexShrink: 0,
                    flexGrow: 0,
                  }}
                />
                <UserAvatarDetails creator={threadToDisplay.creator} />
              </Box>
              <Box
                aria-label={"Thread Created Time"}
                sx={{
                  display: "inline-flex",
                  gap: { xs: "0.4rem", sm: "0.5rem" },
                  alignItems: "center",
                }}
              >
                <UserContentTimestamp
                  createdTimestamp={new Date(threadToDisplay.created_time)}
                  updatedTimestamp={new Date(threadToDisplay.updated_time)}
                />
              </Box>
            </Box>
            {threadToDisplay.creator === auth.username && (
              <>
                <Divider />
                <Box className={"my-4 flex items-center justify-end"}>
                  <Link to={`/viewthread/${threadToDisplay.id}/edit`}>
                    <Button
                      variant="outlined"
                      size="large"
                      sx={{ mr: 2 }}
                    >
                      Edit
                    </Button>
                  </Link>
                  <Button
                    variant="outlined"
                    size="large"
                    color="error"
                    sx={{ mr: 2 }}
                  >
                    Delete
                  </Button>
                </Box>
              </>
            )}
          </Box>

          <Divider sx={{ mx: 4, my: 4 }} />
          <CommentTextField
            setCommentContent={setCommentContent}
            handleSubmit={() => console.log("submit Comment", commentContent)}
          />
          <Divider sx={{ mx: 4, mt: 2, mb: 4 }} />
          <Box
            sx={{
              mx: 4,
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
            }}
          >
            <Typography variant="h5">99 Comments</Typography>
            <SortButton
              availableSortCriteriaMappings={availableCommentSortCriteria}
              setSortCriteria={setCommentSortCriteria}
              size={"large"}
            />
          </Box>
          <Box sx={{ mx: 4, mb: 20 }}>
            {sampleComments.map((comment) => (
              <ThreadComment
                key={comment.id}
                comment={comment}
              />
            ))}
          </Box>
        </>
      )}
    </Box>
  );
}
