/**
 * Standard API response wrapper from Go backend.
 */
export interface ApiResponse<T> {
  data: T;
  message: string;
  success: boolean;
}

/**
 * Paginated API response from Go backend.
 */
export interface PaginatedResponse<T> {
  data: T[];
  message: string;
  success: boolean;
  meta: {
    page: number;
    per_page: number;
    total: number;
    total_pages: number;
  };
}
