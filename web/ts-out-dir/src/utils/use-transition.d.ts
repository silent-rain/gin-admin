import { type Ref } from 'vue';
export declare const TransitionPresets: Record<string, (t: number) => number>;
export interface UseTransitionOptions {
    duration?: number;
    transition?: ((t: number) => number) | keyof typeof TransitionPresets;
    initialValue?: number;
}
export declare function useTransition(source: Ref<number>, options?: UseTransitionOptions): Ref<number>;
