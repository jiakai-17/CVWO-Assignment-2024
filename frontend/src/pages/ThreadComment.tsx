import Box from "@mui/material/Box";
import UserAvatarDetails from "../components/UserAvatarDetails.tsx";
import UserContentTimestamp from "../components/UserContentTimestamp.tsx";
import Typography from "@mui/material/Typography";
import { Divider } from "@mui/material";
import Button from "@mui/material/Button";
import Comment from "../models/comment/Comment.tsx";
import { useState } from "react";
import CommentTextField from "../components/CommentTextField.tsx";

export function ThreadComment(props: Readonly<{ comment: Comment }>) {
  const [commentBody, setCommentBody] = useState(props.comment.body);
  const [isEditing, setIsEditing] = useState(false);

  function handleEdit() {
    setIsEditing(true);
  }

  function handleSave() {
    setIsEditing(false);
  }

  function handleCancel() {
    setIsEditing(false);
  }

  return (
    <Box
      sx={{
        my: 4,
        p: 2,
        pb: 0,
        border: 2,
        borderColor: "lightgray",
        borderRadius: "0.375rem",
      }}
    >
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
          mb: 2,
          maxWidth: "100%",
        }}
      >
        <Box
          sx={{
            display: "flex",
            gap: { xs: "0.4rem", sm: "0.5rem" },
            alignItems: "center",
            pr: 1,
            minWidth: 0,
          }}
        >
          <UserAvatarDetails creator={props.comment.creator} />
        </Box>
        <Box
          aria-label={"Comment Created Time"}
          sx={{
            display: "inline-flex",
            gap: { xs: "0.4rem", sm: "0.5rem" },
            alignItems: "center",
            flexShrink: 0,
          }}
        >
          <UserContentTimestamp
            createdTimestamp={props.comment.created_time}
            updatedTimestamp={props.comment.updated_time}
          />
        </Box>
      </Box>
      {isEditing && (
        <CommentTextField
          setCommentContent={setCommentBody}
          handleSubmit={handleSave}
          defaultContent={commentBody}
          submitButtonLabel={"Save"}
          textFieldLabel={"Edit your comment..."}
          handleCancel={handleCancel}
        />
      )}
      {!isEditing && <Typography sx={{ pb: 2 }}>{props.comment.body}</Typography>}
      {!isEditing && props.comment.creator === localStorage.getItem("username") && (
        <>
          <Divider sx={{ my: 2 }} />
          <Box className={"my-4 flex items-center justify-end"}>
            <Button
              variant="outlined"
              size="large"
              sx={{ mr: 2 }}
              onClick={handleEdit}
            >
              Edit
            </Button>
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
  );
}
