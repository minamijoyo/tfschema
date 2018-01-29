package command

import (
	"github.com/minamijoyo/tfschema/tfschema"
	"github.com/posener/complete"
)

func (m *Meta) completePredictResourceType() complete.Predictor {
	return complete.PredictFunc(func(args complete.Args) []string {

		resourceType := args.Last
		providerName, err := detectProviderName(resourceType)
		if err != nil {
			return nil
		}

		client, err := tfschema.NewClient(providerName)
		if err != nil {
			return nil
		}

		defer client.Kill()

		res := client.Resources()

		resourceTypes := []string{}
		for _, r := range res {
			resourceTypes = append(resourceTypes, r.Name)
		}

		return resourceTypes

	})
}

func (m *Meta) completePredictDataSource() complete.Predictor {
	return complete.PredictFunc(func(args complete.Args) []string {

		dataSource := args.Last
		providerName, err := detectProviderName(dataSource)
		if err != nil {
			return nil
		}

		client, err := tfschema.NewClient(providerName)
		if err != nil {
			return nil
		}

		defer client.Kill()

		res := client.DataSources()

		dataSources := []string{}
		for _, r := range res {
			dataSources = append(dataSources, r.Name)
		}

		return dataSources

	})
}
