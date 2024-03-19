export class ApiError extends Error {
  constructor(
    message: string,
    private status: number,
  ) {
    super(message)
  }

  getStatusCode() {
    return this.status
  }
}
