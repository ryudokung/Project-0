package combat

import (
	"os"
	"gopkg.in/yaml.v3"
)

type BalanceConfig struct {
	Resonance struct {
		GainRateDealt            float64 `yaml:"gain_rate_dealt"`
		GainRateTaken            float64 `yaml:"gain_rate_taken"`
		BonusAccuracyPerLevel    int     `yaml:"bonus_accuracy_per_level"`
		BonusEvasionPerLevel     int     `yaml:"bonus_evasion_per_level"`
		ResonanceDamageMultiplier float64 `yaml:"resonance_damage_multiplier"`
	} `yaml:"resonance"`

	Progression struct {
		BaseSyncRate     float64 `yaml:"base_sync_rate"`
		SyncRatePerLevel float64 `yaml:"sync_rate_per_level"`
	} `yaml:"progression"`

	ScaleSuppression struct {
		HumanVsMechDamageReduction        float64 `yaml:"human_vs_mech_damage_reduction"`
		MechVsHumanDamageMultiplier       float64 `yaml:"mech_vs_human_damage_multiplier"`
		ResonantHumanVsMechDamageMultiplier float64 `yaml:"resonant_human_vs_mech_damage_multiplier"`
		ResonantHumanDeflectionRate       float64 `yaml:"resonant_human_deflection_rate"`
	} `yaml:"scale_suppression"`

	BaseStats struct {
		DefaultAccuracy         int     `yaml:"default_accuracy"`
		DefaultDefenseEfficiency float64 `yaml:"default_defense_efficiency"`
		ForcedSurvivalHP        int     `yaml:"forced_survival_hp"`
		BossPhase2HP            int     `yaml:"boss_phase_2_hp"`
		BossPhase2Attack        int     `yaml:"boss_phase_2_attack"`
		BossPhase2ResonanceLevel int     `yaml:"boss_phase_2_resonance_level"`
	} `yaml:"base_stats"`
}

var GlobalBalance BalanceConfig

func LoadBalanceConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &GlobalBalance)
}
