import * as React from "react";
import Button from "@mui/material/Button";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import SortIcon from "@mui/icons-material/Sort";
import { useState } from "react";

// Creates a button that opens a menu to show available sort criteria
export default function SortButton(
  props: Readonly<{
    availableSortCriteriaMappings: Map<string, string>;
    setSortCriteria: (sortCriteria: string) => void;
    size: "small" | "large";
  }>,
) {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);

  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleMenuItemClick = (criteria: string) => {
    props.setSortCriteria(criteria);
    handleClose();
  };

  return (
    <>
      <Button
        id="sort-button"
        aria-controls={open ? "basic-menu" : undefined}
        aria-haspopup="true"
        aria-expanded={open ? "true" : undefined}
        onClick={handleClick}
        variant="outlined"
        size={props.size}
        endIcon={props.size === "large" ? <SortIcon /> : undefined}
      >
        {props.size === "large" ? "Sort" : <SortIcon />}
      </Button>
      <Menu
        id="basic-menu"
        anchorEl={anchorEl}
        open={open}
        onClose={handleClose}
        MenuListProps={{
          "aria-labelledby": "sort-button",
        }}
      >
        {Array.from(props.availableSortCriteriaMappings.keys()).map((criteria) => (
          <MenuItem
            key={criteria}
            onClick={() => handleMenuItemClick(criteria)}
          >
            {criteria}
          </MenuItem>
        ))}
      </Menu>
    </>
  );
}
