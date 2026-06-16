import { ref, watch, onUnmounted } from 'vue';
export const TransitionPresets = {
    linear: (t) => t,
    easeInSine: (t) => 1 - Math.cos((t * Math.PI) / 2),
    easeOutSine: (t) => Math.sin((t * Math.PI) / 2),
    easeInOutSine: (t) => -(Math.cos(Math.PI * t) - 1) / 2,
    easeInQuad: (t) => t * t,
    easeOutQuad: (t) => 1 - (1 - t) * (1 - t),
    easeInOutQuad: (t) => t < 0.5 ? 2 * t * t : 1 - Math.pow(-2 * t + 2, 2) / 2,
    easeInCubic: (t) => t * t * t,
    easeOutCubic: (t) => 1 - Math.pow(1 - t, 3),
    easeInOutCubic: (t) => t < 0.5 ? 4 * t * t * t : 1 - Math.pow(-2 * t + 2, 3) / 2,
    easeInQuart: (t) => t * t * t * t,
    easeOutQuart: (t) => 1 - Math.pow(1 - t, 4),
    easeInOutQuart: (t) => t < 0.5 ? 8 * t * t * t * t : 1 - Math.pow(-2 * t + 2, 4) / 2,
    easeInQuint: (t) => t * t * t * t * t,
    easeOutQuint: (t) => 1 - Math.pow(1 - t, 5),
    easeInOutQuint: (t) => t < 0.5 ? 16 * t * t * t * t * t : 1 - Math.pow(-2 * t + 2, 5) / 2,
    easeInExpo: (t) => (t === 0 ? 0 : Math.pow(2, 10 * t - 10)),
    easeOutExpo: (t) => (t === 1 ? 1 : 1 - Math.pow(2, -10 * t)),
    easeInOutExpo: (t) => {
        if (t === 0)
            return 0;
        if (t === 1)
            return 1;
        return t < 0.5
            ? Math.pow(2, 20 * t - 10) / 2
            : (2 - Math.pow(2, -20 * t + 10)) / 2;
    },
    easeInCirc: (t) => 1 - Math.sqrt(1 - Math.pow(t, 2)),
    easeOutCirc: (t) => Math.sqrt(1 - Math.pow(t - 1, 2)),
    easeInOutCirc: (t) => t < 0.5
        ? (1 - Math.sqrt(1 - Math.pow(2 * t, 2))) / 2
        : (Math.sqrt(1 - Math.pow(-2 * t + 2, 2)) + 1) / 2,
    easeInBack: (t) => {
        const c1 = 1.70158;
        return c1 * t * t * t - c1 * t * t;
    },
    easeOutBack: (t) => {
        const c1 = 1.70158;
        return 1 + c1 * Math.pow(t - 1, 3) + c1 * Math.pow(t - 1, 2);
    },
    easeInOutBack: (t) => {
        const c1 = 1.70158;
        const c2 = c1 * 1.525;
        return t < 0.5
            ? (Math.pow(2 * t, 2) * ((c2 + 1) * 2 * t - c2)) / 2
            : (Math.pow(2 * t - 2, 2) * ((c2 + 1) * (t * 2 - 2) + c2) + 2) / 2;
    },
};
export function useTransition(source, options = {}) {
    const { duration = 1000, transition = TransitionPresets.easeInOutCubic, initialValue = 0, } = options;
    const output = ref(initialValue);
    let animationFrameId = null;
    let startTime = null;
    let startValue = initialValue;
    let targetValue = source.value;
    let isAnimating = false;
    const getEasingFunction = () => {
        if (typeof transition === 'function') {
            return transition;
        }
        return TransitionPresets[transition] || TransitionPresets.easeInOutCubic;
    };
    const animate = (currentTime) => {
        if (startTime === null) {
            startTime = currentTime;
        }
        const elapsed = currentTime - startTime;
        const progress = Math.min(elapsed / duration, 1);
        const easingFunction = getEasingFunction();
        const easedProgress = easingFunction(progress);
        output.value = startValue + (targetValue - startValue) * easedProgress;
        if (progress < 1) {
            animationFrameId = requestAnimationFrame(animate);
        }
        else {
            isAnimating = false;
            animationFrameId = null;
            startTime = null;
        }
    };
    const startAnimation = () => {
        if (isAnimating) {
            if (animationFrameId !== null) {
                cancelAnimationFrame(animationFrameId);
            }
        }
        startValue = output.value;
        targetValue = source.value;
        startTime = null;
        isAnimating = true;
        animationFrameId = requestAnimationFrame(animate);
    };
    const stopAnimation = () => {
        if (animationFrameId !== null) {
            cancelAnimationFrame(animationFrameId);
            animationFrameId = null;
        }
        isAnimating = false;
        startTime = null;
    };
    const stopWatch = watch(() => source.value, () => {
        startAnimation();
    });
    onUnmounted(() => {
        stopAnimation();
        stopWatch();
    });
    return output;
}
