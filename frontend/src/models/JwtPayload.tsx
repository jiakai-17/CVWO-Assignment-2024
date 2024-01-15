export type JWTPayload = {
  username: string;
  iat: number;
  exp: number;
};

export default JWTPayload;
