import TextField from "@mui/material/TextField";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { useContext, useEffect, useState } from "react";
import Typography from "@mui/material/Typography";
import AuthContext from "../contexts/AuthContext.tsx";

// Creates a text field for the user to create/edit a new comment
export default function CommentTextField(
  props: Readonly<{
    setCommentContent: (content: string) => void;
    handleSubmit: () => void;
    defaultContent?: string;
    submitButtonLabel?: string;
    textFieldLabel?: string;
    handleCancel?: () => void;
  }>,
) {
  const { auth, isLoaded } = useContext(AuthContext);

  const [isLogin, setIsLogin] = useState(auth.isLogin);

  useEffect(() => {
    if (!isLoaded) {
      return;
    }
    setIsLogin(auth.isLogin);
  }, [auth, isLoaded]);

  const [commentContent, setCommentContent] = useState(props.defaultContent ?? "");

  useEffect(() => {
    setCommentContent(props.defaultContent ?? "");
  }, [props.defaultContent]);

  const setCommentCallback = props.setCommentContent;
  const handleSubmitCallback = props.handleSubmit;

  useEffect(() => {
    setCommentCallback(commentContent);
  }, [commentContent, setCommentCallback]);

  function onSubmit() {
    handleSubmitCallback();
  }

  return (
    <Box sx={{ mx: 4, width: "full" }}>
      {!isLogin && (
        <Typography sx={{ mb: 4, textAlign: "center", color: "gray" }}>
          You need to log in to leave a comment.
        </Typography>
      )}
      {isLogin && (
        <>
          <TextField
            id="outlined-multiline-static"
            label={props.textFieldLabel ?? "Leave a comment... (max 3000 chars)"}
            multiline
            rows={3}
            variant="outlined"
            value={commentContent}
            onChange={(event) => setCommentContent(event.target.value)}
            fullWidth
          />
          <Box className={"flex justify-end gap-3"}>
            {props.handleCancel && (
              <Button
                variant="outlined"
                size={"large"}
                sx={{ my: 2 }}
                onClick={props.handleCancel}
                color={"error"}
              >
                Cancel
              </Button>
            )}
            <Button
              variant="contained"
              size={"large"}
              sx={{ my: 2 }}
              onClick={onSubmit}
            >
              {props.submitButtonLabel ?? "Submit"}
            </Button>
          </Box>
        </>
      )}
    </Box>
  );
}
