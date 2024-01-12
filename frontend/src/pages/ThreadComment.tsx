import Box from "@mui/material/Box";
import UserAvatarDetails from "../components/UserAvatarDetails.tsx";
import UserContentTimestamp from "../components/UserContentTimestamp.tsx";
import Typography from "@mui/material/Typography";
import { Divider } from "@mui/material";
import Button from "@mui/material/Button";
import Comment from "../models/comment/Comment.tsx";
import { useContext, useEffect, useState } from "react";
import CommentTextField from "../components/CommentTextField.tsx";
import AuthContext from "../contexts/AuthContext.tsx";

export function ThreadComment(
  props: Readonly<{
    comment: Comment;
    deleteComment: () => void;
  }>,
) {
  const [currentComment, setCurrentComment] = useState(props.comment);
  const [editedCommentBody, setEditedCommentBody] = useState(props.comment.body);

  const [isEditing, setIsEditing] = useState(false);

  const { auth, isLoaded } = useContext(AuthContext);

  const [isCommentCreator, setIsCommentCreator] = useState(auth.username === props.comment.creator);

  useEffect(() => {
    if (!isLoaded) {
      return;
    }
    setIsCommentCreator(auth.username === props.comment.creator);
  }, [auth, isLoaded, props.comment.creator]);

  function handleEdit() {
    setIsEditing(true);
  }

  function handleSave() {
    fetch(`/api/v1/comment/${props.comment.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${auth.token}`,
      },
      body: JSON.stringify({ body: editedCommentBody }),
    }).then((res) => {
      if (!res.ok) {
        console.log("edit failed");
        res.text().then((text) => console.log(text));
      } else {
        setIsEditing(false);
        setCurrentComment({
          ...currentComment,
          body: editedCommentBody,
          updated_time: new Date().toISOString(),
        });
      }
    });
  }

  function handleCancel() {
    setIsEditing(false);
  }

  // Delete comment
  const [isDeleted, setIsDeleted] = useState(false);
  const handleDelete = () => {
    const shouldDelete = window.confirm("Are you sure you want to delete this comment? This action cannot be undone.");
    if (!shouldDelete) {
      return;
    } else {
      fetch(`/api/v1/comment/${props.comment.id}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${auth.token}`,
        },
      }).then((res) => {
        if (!res.ok) {
          console.log("delete failed");
          res.text().then((text) => console.log(text));
        } else {
          console.log("delete success");
          setIsDeleted(true);
          props.deleteComment();
        }
      });
    }
  };

  return (
    <>
      {!isDeleted && (
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
              <UserAvatarDetails creator={currentComment.creator} />
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
                createdTimestamp={new Date(currentComment.created_time)}
                updatedTimestamp={new Date(currentComment.updated_time)}
              />
            </Box>
          </Box>
          {isEditing && (
            <CommentTextField
              setCommentContent={setEditedCommentBody}
              handleSubmit={handleSave}
              defaultContent={currentComment.body}
              submitButtonLabel={"Save"}
              textFieldLabel={"Edit your comment..."}
              handleCancel={handleCancel}
            />
          )}
          {!isEditing && <Typography sx={{ pb: 2, whiteSpace: "pre-line" }}>{currentComment.body}</Typography>}
          {!isEditing && isCommentCreator && (
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
                  onClick={handleDelete}
                >
                  Delete
                </Button>
              </Box>
            </>
          )}
        </Box>
      )}
    </>
  );
}
