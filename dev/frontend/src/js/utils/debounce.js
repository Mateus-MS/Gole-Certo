/**
 * Adds debounce behavior to a function.
 *
 * Debounce: ensures a function is called only after a specified delay
 * has passed since the last time it was invoked.
 *
 * @param {Function} func - The function to debounce.
 * @param {number} delay - Delay in milliseconds before invoking `func`.
 * @returns {Function} A new debounced version of `func`.
 *
 * @example
 * const log = debounce(() => console.log('Fired!'), 500);
 * window.addEventListener('resize', log);
 * // Logs "Fired!" only once after the user stops resizing for 500ms.
 *
 * @file "/js/utils/debounce.js"
 */
function debounce(func, delay) {
    let timeoutId;
    return function (...args) {
        clearTimeout(timeoutId);
        timeoutId = setTimeout(() => {
            func.apply(this, args);
        }, delay);
    };
}