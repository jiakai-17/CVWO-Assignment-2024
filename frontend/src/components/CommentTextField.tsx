import TextField from "@mui/material/TextField";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import * as React from "react";
import { useEffect, useState } from "react";
import Typography from "@mui/material/Typography";

export default function CommentTextField(props: Readonly<{
  setCommentContent: (content: string) => void
  handleSubmit: () => void
}>) {

  const isLogin = localStorage.getItem("isLogin") === "true";

  const [commentContent, setCommentContent] = useState("");

  const setCommentCallback = props.setCommentContent;
  const handleSubmitCallback = props.handleSubmit;

  useEffect(() => {
    setCommentCallback(commentContent);
  }, [commentContent, setCommentCallback]);

  function onSubmit() {
    handleSubmitCallback();
    setCommentContent("");
  }

  return (
    <Box sx={{ mx: 4, width: "full" }}>
      {!isLogin &&
        <Typography sx={{ mb: 4, textAlign: "center", color: "gray" }}>
          You need to log in to leave a comment.
        </Typography>
      }
      {isLogin &&
        <>
          <TextField
            id="outlined-multiline-static"
            label="Leave a comment..."
            multiline
            rows={3}
            variant="outlined"
            value={commentContent}
            onChange={(event) => setCommentContent(event.target.value)}
            fullWidth
          />
          <Box sx={{ display: "flex", justifyContent: "flex-end" }}>
            <Button
              variant="contained"
              size={"large"}
              sx={{ my: 2 }}
              onClick={onSubmit}
            >
              Comment
            </Button>
          </Box>
        </>
      }
    </Box>
  );
}
