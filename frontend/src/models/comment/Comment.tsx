// Defines a Comment type for the application.
export type Comment = {
  id: string;
  body: string;
  creator: string;
  thread_id: string;
  created_time: string;
  updated_time: string;
};

export default Comment;
