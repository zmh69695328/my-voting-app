// Check if a value is not null, undefined, or an empty string
export function isNotEmpty(value) {
  return value !== null && value !== undefined && value !== '';
}
