package main

import (
	"context"

	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/mgutz/ansi"
	"github.com/pkg/errors"
)

// updateAccessControl apply accesses to a dataset according to conf
func updateAccessControl(client *bigquery.Client, conf Config) error {

	for _, dataset := range conf.Datasets {

		fmt.Println(ansi.Color(fmt.Sprintf("\n\nUpdating accesses for %s:%s", conf.Project, dataset.Name), "blue+b"))

		metaToUpdate := &bigquery.DatasetMetadataToUpdate{}

		generateAccesses(metaToUpdate, dataset.Owner.GroupByEmail, bigquery.OwnerRole, bigquery.GroupEmailEntity)
		generateAccesses(metaToUpdate, dataset.Owner.UserByEmail, bigquery.OwnerRole, bigquery.UserEmailEntity)
		generateAccesses(metaToUpdate, dataset.Owner.SpecialGroup, bigquery.OwnerRole, bigquery.SpecialGroupEntity)
		generateAccesses(metaToUpdate, dataset.Writer.GroupByEmail, bigquery.WriterRole, bigquery.GroupEmailEntity)
		generateAccesses(metaToUpdate, dataset.Writer.UserByEmail, bigquery.WriterRole, bigquery.UserEmailEntity)
		generateAccesses(metaToUpdate, dataset.Writer.SpecialGroup, bigquery.WriterRole, bigquery.SpecialGroupEntity)
		generateAccesses(metaToUpdate, dataset.Reader.GroupByEmail, bigquery.ReaderRole, bigquery.GroupEmailEntity)
		generateAccesses(metaToUpdate, dataset.Reader.UserByEmail, bigquery.ReaderRole, bigquery.UserEmailEntity)
		generateAccesses(metaToUpdate, dataset.Reader.SpecialGroup, bigquery.ReaderRole, bigquery.SpecialGroupEntity)

		// TODO: Factorize with above
		for _, view := range dataset.View {
			metaToUpdate.Access = append(metaToUpdate.Access, &bigquery.AccessEntry{
				EntityType: bigquery.ViewEntity,
				View: &bigquery.Table{
					ProjectID: conf.Project,
					DatasetID: view.DatasetID,
					TableID:   view.ViewID,
				},
			})
		}

		ds := client.Dataset(dataset.Name)
		ctx := context.Background()
		meta, err := ds.Metadata(ctx)
		if err != nil {
			return errors.Wrap(err, "cannot get metadata for original dataset")
		}

		// Print the diff
		isDiff := diff(AccessList(meta.Access), AccessList(metaToUpdate.Access))

		// Update access
		if isDiff {
			if _, err := ds.Update(ctx, *metaToUpdate, meta.ETag); err != nil {
				return errors.Wrap(err, "cannot update original dataset's metadata")
			}
		}
	}

	return nil
}

// generateAccesses update the AccessEntries for a given bigquery.DatasetMetadataToUpdate
func generateAccesses(meta *bigquery.DatasetMetadataToUpdate, entities []string, role bigquery.AccessRole,
	entityType bigquery.EntityType) {

	for _, entity := range entities {
		meta.Access = append(meta.Access, &bigquery.AccessEntry{
			Role:       role,
			EntityType: entityType,
			Entity:     entity,
		})
	}
}
