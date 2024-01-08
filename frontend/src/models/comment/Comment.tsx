// Defines a Comment type for the application.
export type Comment = {
  id: string;
  body: string;
  creator: string;
  thread_id: string;
  created_time: Date;
  updated_time: Date;
}

export default Comment;
