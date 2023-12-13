"use client";

import * as React from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { Divider } from "@mui/material";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import PublishIcon from "@mui/icons-material/Publish";
import Link from "next/link";

export default function Page() {

  return (
    <Box sx={{ mx: "2", textAlign: "center", mb: 10 }}>
      <Typography variant="h4" sx={{ px: 2 }}>Create a new Thread</Typography>
      <Divider sx={{ m: 5 }} />
      <Box
        sx={{ mx: { xs: 2, sm: 4 }, display: "flex", flexWrap: "wrap", flexDirection: "column" }}>
        <TextField
          id="outlined-basic"
          label="Title"
          variant="outlined"
          sx={{ my: 2, flexGrow: 1 }}
        />
        <TextField
          id="outlined-multiline-static"
          label="Body"
          multiline
          rows={10}
          variant="outlined"
          sx={{ my: 2, flexGrow: 1 }}
        />
        <Box sx={{ display: "flex", justifyContent: "space-between" }}>
          <Link href={"/"}>
            <Button
              variant="outlined"
              size="large"
              color="error"
              sx={{ mr: 2 }}>
              Cancel
            </Button>
          </Link>
          <Button
            variant="contained"
            disableElevation
            size="large"
            endIcon={<PublishIcon />}
            sx={{ width: "12rem" }}>
            Post
          </Button>
        </Box>
      </Box>
    </Box>
  );
}
