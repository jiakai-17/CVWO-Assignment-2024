import React from "react";
import { toSvg } from "jdenticon";
import Image from "next/image";

export default function AvatarIdenticon({ username }: Readonly<{
  username: string,
  size?: number
}>) {
  const svg = toSvg(username, 50);
  return (
    <Image alt={username}
           src={`data:image/svg+xml;utf8,${encodeURIComponent(svg)}`}
           width={50}
           height={50} />
  );
}
