import type { Ref } from 'vue'
import { onUnmounted, ref, watch } from 'vue'

/**
 * Transition presets for easing functions
 * Ported from @vueuse/core to avoid Rolldown INVALID_ANNOTATION warnings
 */
export const TransitionPresets: Record<string, (t: number) => number> = {
  linear: (t: number) => t,
  easeInSine: (t: number) => 1 - Math.cos((t * Math.PI) / 2),
  easeOutSine: (t: number) => Math.sin((t * Math.PI) / 2),
  easeInOutSine: (t: number) => -(Math.cos(Math.PI * t) - 1) / 2,
  easeInQuad: (t: number) => t * t,
  easeOutQuad: (t: number) => 1 - (1 - t) * (1 - t),
  easeInOutQuad: (t: number) =>
    t < 0.5 ? 2 * t * t : 1 - (-2 * t + 2) ** 2 / 2,
  easeInCubic: (t: number) => t * t * t,
  easeOutCubic: (t: number) => 1 - (1 - t) ** 3,
  easeInOutCubic: (t: number) =>
    t < 0.5 ? 4 * t * t * t : 1 - (-2 * t + 2) ** 3 / 2,
  easeInQuart: (t: number) => t * t * t * t,
  easeOutQuart: (t: number) => 1 - (1 - t) ** 4,
  easeInOutQuart: (t: number) =>
    t < 0.5 ? 8 * t * t * t * t : 1 - (-2 * t + 2) ** 4 / 2,
  easeInQuint: (t: number) => t * t * t * t * t,
  easeOutQuint: (t: number) => 1 - (1 - t) ** 5,
  easeInOutQuint: (t: number) =>
    t < 0.5 ? 16 * t * t * t * t * t : 1 - (-2 * t + 2) ** 5 / 2,
  easeInExpo: (t: number) => (t === 0 ? 0 : 2 ** (10 * t - 10)),
  easeOutExpo: (t: number) => (t === 1 ? 1 : 1 - 2 ** (-10 * t)),
  easeInOutExpo: (t: number) => {
    if (t === 0)
      return 0
    if (t === 1)
      return 1
    return t < 0.5
      ? 2 ** (20 * t - 10) / 2
      : (2 - 2 ** (-20 * t + 10)) / 2
  },
  easeInCirc: (t: number) => 1 - Math.sqrt(1 - t ** 2),
  easeOutCirc: (t: number) => Math.sqrt(1 - (t - 1) ** 2),
  easeInOutCirc: (t: number) =>
    t < 0.5
      ? (1 - Math.sqrt(1 - (2 * t) ** 2)) / 2
      : (Math.sqrt(1 - (-2 * t + 2) ** 2) + 1) / 2,
  easeInBack: (t: number) => {
    const c1 = 1.70158
    return c1 * t * t * t - c1 * t * t
  },
  easeOutBack: (t: number) => {
    const c1 = 1.70158
    return 1 + c1 * (t - 1) ** 3 + c1 * (t - 1) ** 2
  },
  easeInOutBack: (t: number) => {
    const c1 = 1.70158
    const c2 = c1 * 1.525
    return t < 0.5
      ? ((2 * t) ** 2 * ((c2 + 1) * 2 * t - c2)) / 2
      : ((2 * t - 2) ** 2 * ((c2 + 1) * (t * 2 - 2) + c2) + 2) / 2
  },
}

export interface UseTransitionOptions {
  /**
   * Transition duration in milliseconds
   * @default 1000
   */
  duration?: number
  /**
   * Easing function or preset name
   * @default TransitionPresets.easeInOutCubic
   */
  transition?: ((t: number) => number) | keyof typeof TransitionPresets
  /**
   * Initial value
   * @default 0
   */
  initialValue?: number
}

/**
 * Reactive value that transitions smoothly when the source value changes
 * @param source - The source ref to transition from
 * @param options - Transition options
 * @returns A ref with the transitioning value
 */
export function useTransition(
  source: Ref<number>,
  options: UseTransitionOptions = {},
): Ref<number> {
  const {
    duration = 1000,
    transition = TransitionPresets.easeInOutCubic,
    initialValue = 0,
  } = options

  const output = ref<number>(initialValue) as Ref<number>

  let animationFrameId: number | null = null
  let startTime: number | null = null
  let startValue: number = initialValue
  let targetValue: number = source.value
  let isAnimating = false

  const getEasingFunction = (): ((t: number) => number) => {
    if (typeof transition === 'function') {
      return transition
    }
    return TransitionPresets[transition] || TransitionPresets.easeInOutCubic
  }

  const animate = (currentTime: number) => {
    if (startTime === null) {
      startTime = currentTime
    }

    const elapsed = currentTime - startTime
    const progress = Math.min(elapsed / duration, 1)
    const easingFunction = getEasingFunction()
    const easedProgress = easingFunction(progress)

    output.value = startValue + (targetValue - startValue) * easedProgress

    if (progress < 1) {
      animationFrameId = requestAnimationFrame(animate)
    }
    else {
      isAnimating = false
      animationFrameId = null
      startTime = null
    }
  }

  const startAnimation = () => {
    if (isAnimating) {
      // Cancel current animation and start new one
      if (animationFrameId !== null) {
        cancelAnimationFrame(animationFrameId)
      }
    }

    startValue = output.value
    targetValue = source.value
    startTime = null
    isAnimating = true
    animationFrameId = requestAnimationFrame(animate)
  }

  const stopAnimation = () => {
    if (animationFrameId !== null) {
      cancelAnimationFrame(animationFrameId)
      animationFrameId = null
    }
    isAnimating = false
    startTime = null
  }

  // Watch for source value changes
  const stopWatch = watch(
    () => source.value,
    () => {
      startAnimation()
    },
  )

  // Cleanup on unmount
  onUnmounted(() => {
    stopAnimation()
    stopWatch()
  })

  return output
}
