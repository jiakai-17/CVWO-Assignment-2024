"use client";

import { Chip } from "@mui/material";
import { Link, useSearchParams } from "react-router-dom";
import { useEffect, useState } from "react";

// Creates a tag component to be displayed in the list of tags on the main page
export default function ThreadTag(props: Readonly<{ tag: string }>) {
  const [searchParams] = useSearchParams();
  const [hasSortCriteria, setHasSortCriteria] = useState(searchParams.has("order"));
  const [sortCriteria, setSortCriteria] = useState(searchParams.get("order") ?? "Newest first");

  useEffect(() => {
    setHasSortCriteria(searchParams.has("order"));
  }, [searchParams]);

  useEffect(() => {
    if (hasSortCriteria) {
      setSortCriteria(searchParams.get("order") ?? "Newest first");
    }
  }, [hasSortCriteria, searchParams]);

  return (
    <Link to={hasSortCriteria ? `/?q=tag:${props.tag}&order=${sortCriteria}` : `/?q=tag:${props.tag}`}>
      <Chip
        key={props.tag}
        label={props.tag}
        onClick={() => {
          // This empty function enables the hover effect on the chip
        }}
      />
    </Link>
  );
}
