"use client";

import { Chip } from "@mui/material";
import { Link } from "react-router-dom";

export default function ThreadTag(props: Readonly<{ tag: string }>) {
  const handleChipClick = (tag: string) => {
    console.info(`You clicked on the ${tag} tag.`);
  };

  return (
    <Link to={`/?q=tag:${props.tag}`}>
      <Chip
        key={props.tag}
        label={props.tag}
        onClick={() => handleChipClick(props.tag)}
      />
    </Link>
  );
}
